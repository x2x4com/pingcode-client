package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WikiSpace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Identifier  string `json:"identifier"`
	ScopeType   string `json:"scope_type"`
	ScopeID     string `json:"scope_id,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
	Description string `json:"description,omitempty"`
}

type WikiSpaceMember struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type WikiSpaceMemberListResponse struct {
	Values    []WikiSpaceMember `json:"values"`
	Total     int               `json:"total"`
	PageSize  int               `json:"page_size"`
	PageIndex int               `json:"page_index"`
}

type WikiPage struct {
	ID       string `json:"id"`
	Title    string `json:"name"`
	SpaceID  string `json:"space_id,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
	Type     string `json:"type,omitempty"`
}

type WikiPageContent struct {
	PageID  string `json:"page_id"`
	Content string `json:"content"`
	Format  string `json:"format,omitempty"`
}

type WikiPageVersion struct {
	ID        string `json:"id"`
	Version   int    `json:"version"`
	CreatedAt int64  `json:"created_at"`
}

type WikiSpaceListResponse struct {
	Values    []WikiSpace `json:"values"`
	Total     int         `json:"total"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
}

type WikiPageListResponse struct {
	Values    []WikiPage `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type WikiPageVersionListResponse struct {
	Values    []WikiPageVersion `json:"values"`
	Total     int               `json:"total"`
	PageSize  int               `json:"page_size"`
	PageIndex int               `json:"page_index"`
}

// Space operations

func (c *Client) ListWikiSpaces(pageIndex, pageSize int) ([]WikiSpace, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/spaces?page_index=%d&page_size=%d", pageIndex, pageSize), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list wiki spaces failed with status: %d", resp.StatusCode)
	}
	var listResp WikiSpaceListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) GetWikiSpace(spaceID string) (*WikiSpace, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/spaces/%s", spaceID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get wiki space failed with status: %d", resp.StatusCode)
	}
	var space WikiSpace
	if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
		return nil, err
	}
	return &space, nil
}

func (c *Client) CreateWikiSpace(name, identifier, scopeType, scopeID, visibility, description string) (*WikiSpace, error) {
	body := map[string]interface{}{
		"name":       name,
		"identifier": identifier,
		"scope_type": scopeType,
	}
	if scopeID != "" {
		body["scope_id"] = scopeID
	}
	if visibility != "" {
		body["visibility"] = visibility
	}
	if description != "" {
		body["description"] = description
	}
	resp, err := c.Request("POST", "/wiki/spaces", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create wiki space failed with status: %d", resp.StatusCode)
	}
	var space WikiSpace
	if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
		return nil, err
	}
	return &space, nil
}

func (c *Client) UpdateWikiSpace(spaceID, name, visibility, description string) (*WikiSpace, error) {
	body := map[string]interface{}{}
	if name != "" {
		body["name"] = name
	}
	if visibility != "" {
		body["visibility"] = visibility
	}
	if description != "" {
		body["description"] = description
	}
	resp, err := c.Request("PATCH", fmt.Sprintf("/wiki/spaces/%s", spaceID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update wiki space failed with status: %d", resp.StatusCode)
	}
	var space WikiSpace
	if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
		return nil, err
	}
	return &space, nil
}

func (c *Client) DeleteWikiSpace(spaceID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/wiki/spaces/%s", spaceID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete wiki space failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Space member operations

func (c *Client) ListWikiSpaceMembers(spaceID string) ([]WikiSpaceMember, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/spaces/%s/members", spaceID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list wiki space members failed with status: %d", resp.StatusCode)
	}
	var listResp WikiSpaceMemberListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) AddWikiSpaceMember(spaceID, memberID, memberType, roleID string) (*WikiSpaceMember, error) {
	body := map[string]interface{}{
		"member": map[string]interface{}{
			"id":   memberID,
			"type": memberType,
		},
	}
	if roleID != "" {
		body["role_id"] = roleID
	}
	resp, err := c.Request("POST", fmt.Sprintf("/wiki/spaces/%s/members", spaceID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("add wiki space member failed with status: %d", resp.StatusCode)
	}
	var member WikiSpaceMember
	if err := json.NewDecoder(resp.Body).Decode(&member); err != nil {
		return nil, err
	}
	return &member, nil
}

func (c *Client) RemoveWikiSpaceMember(spaceID, memberID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/wiki/spaces/%s/members/%s", spaceID, memberID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("remove wiki space member failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Page operations

func (c *Client) ListWikiPages(spaceID string, pageIndex, pageSize int) ([]WikiPage, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/pages?space_id=%s&page_index=%d&page_size=%d", spaceID, pageIndex, pageSize), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list wiki pages failed with status: %d", resp.StatusCode)
	}
	var listResp WikiPageListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) GetWikiPage(pageID string) (*WikiPage, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/pages/%s", pageID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get wiki page failed with status: %d", resp.StatusCode)
	}
	var page WikiPage
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) CreateWikiPage(spaceID, parentID, title, pageType string) (*WikiPage, error) {
	body := map[string]interface{}{
		"space_id": spaceID,
		"name":     title,
	}
	if parentID != "" {
		body["parent_id"] = parentID
	}
	if pageType != "" {
		body["type"] = pageType
	}
	resp, err := c.Request("POST", "/wiki/pages", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create wiki page failed with status: %d", resp.StatusCode)
	}
	var page WikiPage
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) UpdateWikiPage(pageID, title string) (*WikiPage, error) {
	body := map[string]interface{}{}
	if title != "" {
		body["name"] = title
	}
	resp, err := c.Request("PATCH", fmt.Sprintf("/wiki/pages/%s", pageID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update wiki page failed with status: %d", resp.StatusCode)
	}
	var page WikiPage
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) DeleteWikiPage(pageID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/wiki/pages/%s", pageID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete wiki page failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Page content operations

func (c *Client) GetWikiPageContent(pageID string) (*WikiPageContent, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/pages/%s/content", pageID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get wiki page content failed with status: %d", resp.StatusCode)
	}
	var content WikiPageContent
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, err
	}
	return &content, nil
}

func (c *Client) UpdateWikiPageContent(pageID, content, format string) (*WikiPageContent, error) {
	body := map[string]interface{}{
		"content":     content,
		"format_type": format,
	}
	resp, err := c.Request("PUT", fmt.Sprintf("/wiki/pages/%s/content", pageID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update wiki page content failed with status: %d", resp.StatusCode)
	}
	var pageContent WikiPageContent
	if err := json.NewDecoder(resp.Body).Decode(&pageContent); err != nil {
		return nil, err
	}
	return &pageContent, nil
}

// Page version operations

func (c *Client) ListWikiPageVersions(pageID string) ([]WikiPageVersion, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/pages/%s/versions", pageID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list wiki page versions failed with status: %d", resp.StatusCode)
	}
	var listResp WikiPageVersionListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}
