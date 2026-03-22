package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultBaseURL = "https://open.pingcode.com/v1"
	MaxRetries     = 3
)

type Client struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	HTTPClient   *http.Client
	AccessToken  string
	TokenExpiry  time.Time
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		BaseURL:      DefaultBaseURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) Authenticate() error {
	url := fmt.Sprintf("%s/auth/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		c.BaseURL, c.ClientID, c.ClientSecret)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status: %d", resp.StatusCode)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return err
	}

	c.AccessToken = tokenResp.AccessToken
	// Note: PingCode's expires_in seems to be a timestamp or a long duration.
	// In the example output: "expires_in":1776710626, which looks like a timestamp.
	// But OAuth2 spec usually says it's seconds from now.
	// Given the large value, I'll treat it as a Unix timestamp if it's > 1e9.
	if tokenResp.ExpiresIn > 1e9 {
		c.TokenExpiry = time.Unix(tokenResp.ExpiresIn, 0)
	} else {
		c.TokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	}

	return nil
}

func (c *Client) Request(method, path string, body interface{}) (*http.Response, error) {
	var lastResp *http.Response
	var err error

	for i := 0; i <= MaxRetries; i++ {
		if c.AccessToken == "" || time.Now().After(c.TokenExpiry) {
			if err := c.Authenticate(); err != nil {
				return nil, err
			}
		}

		url := fmt.Sprintf("%s%s", c.BaseURL, path)
		var bodyReader io.Reader
		if body != nil {
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(jsonBody)
		}

		req, err := http.NewRequest(method, url, bodyReader)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		lastResp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}

		if lastResp.StatusCode == http.StatusTooManyRequests && i < MaxRetries {
			retryAfterStr := lastResp.Header.Get("X-RateLimit-Retry-After")
			retryAfter := 1 // default wait 1s
			if retryAfterStr != "" {
				if val, err := strconv.Atoi(retryAfterStr); err == nil {
					retryAfter = val
				}
			}

			lastResp.Body.Close()
			log.Printf("Rate limit exceeded (429), retrying in %d seconds (retry %d/%d)...", retryAfter, i+1, MaxRetries)
			time.Sleep(time.Duration(retryAfter) * time.Second)
			continue
		}

		return lastResp, nil
	}

	return lastResp, err
}
