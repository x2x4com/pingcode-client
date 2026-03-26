package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DirectoryUser struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	DisplayName    string `json:"display_name"`
	Email          string `json:"email,omitempty"`
	Mobile         string `json:"mobile,omitempty"`
	Status         string `json:"status,omitempty"`
	DepartmentID   string `json:"department_id,omitempty"`
	JobID          string `json:"job_id,omitempty"`
	EmployeeNumber string `json:"employee_number,omitempty"`
	Password       string `json:"password,omitempty"`
}

type DirectoryUserListResponse struct {
	Values    []DirectoryUser `json:"values"`
	Total     int             `json:"total"`
	PageSize  int             `json:"page_size"`
	PageIndex int             `json:"page_index"`
}

type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id,omitempty"`
	HeadID   string `json:"head_id,omitempty"`
}

type DepartmentListResponse struct {
	Values    []Department `json:"values"`
	Total     int          `json:"total"`
	PageSize  int          `json:"page_size"`
	PageIndex int          `json:"page_index"`
}

func (c *Client) ListDirectoryUsers(keywords, name, departmentIDs string) ([]DirectoryUser, error) {
	q := url.Values{}
	if keywords != "" {
		q.Set("keywords", keywords)
	}
	if name != "" {
		q.Set("name", name)
	}
	if departmentIDs != "" {
		q.Set("department_ids", departmentIDs)
	}
	path := "/directory/users"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list directory users failed with status: %d", resp.StatusCode)
	}
	var r DirectoryUserListResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Values, nil
}

func (c *Client) CreateDirectoryUser(u *DirectoryUser) (*DirectoryUser, error) {
	resp, err := c.Request("POST", "/directory/users", u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create directory user failed with status: %d", resp.StatusCode)
	}
	var created DirectoryUser
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

func (c *Client) UpdateDirectoryUser(userID string, u *DirectoryUser) (*DirectoryUser, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/directory/users/%s", userID), u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update directory user failed with status: %d", resp.StatusCode)
	}
	var updated DirectoryUser
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (c *Client) ListDepartments() ([]Department, error) {
	resp, err := c.Request("GET", "/directory/departments", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list departments failed with status: %d", resp.StatusCode)
	}
	var r DepartmentListResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Values, nil
}

func (c *Client) CreateDepartment(name, parentID, headID string) (*Department, error) {
	body := map[string]interface{}{"name": name}
	if parentID != "" {
		body["parent_id"] = parentID
	}
	if headID != "" {
		body["head_id"] = headID
	}
	resp, err := c.Request("POST", "/directory/departments", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create department failed with status: %d", resp.StatusCode)
	}
	var d Department
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (c *Client) UpdateDepartment(departmentID, name, parentID, headID string) (*Department, error) {
	body := map[string]interface{}{}
	if name != "" {
		body["name"] = name
	}
	if parentID != "" {
		body["parent_id"] = parentID
	}
	if headID != "" {
		body["head_id"] = headID
	}
	resp, err := c.Request("PATCH", fmt.Sprintf("/directory/departments/%s", departmentID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update department failed with status: %d", resp.StatusCode)
	}
	var d Department
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (c *Client) DeleteDepartment(departmentID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/directory/departments/%s", departmentID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete department failed with status: %d", resp.StatusCode)
	}
	return nil
}
