package sdk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_RetryOn429(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 2 {
			w.Header().Set("X-RateLimit-Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"values": []}`)
	}))
	defer server.Close()

	client := NewClient("id", "secret")
	client.BaseURL = server.URL
	client.AccessToken = "fake-token"
	client.TokenExpiry = time.Now().Add(1 * time.Hour)

	_, err := client.Request("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
}

func TestClient_MaxRetriesExceeded(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.Header().Set("X-RateLimit-Retry-After", "0")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	client := NewClient("id", "secret")
	client.BaseURL = server.URL
	client.AccessToken = "fake-token"
	client.TokenExpiry = time.Now().Add(1 * time.Hour)

	resp, err := client.Request("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Expected no error from Request itself (it returns the last response), got %v", err)
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", resp.StatusCode)
	}

	// MaxRetries is 3, so total attempts should be 1 (initial) + 3 (retries) = 4
	if attempts != 4 {
		t.Errorf("Expected 4 attempts, got %d", attempts)
	}
}
