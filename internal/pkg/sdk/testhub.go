package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TestLibrary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Identifier  string `json:"identifier,omitempty"`
	ScopeType   string `json:"scope_type,omitempty"`
	ScopeID     string `json:"scope_id,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
	Description string `json:"description,omitempty"`
}

type TestLibraryListResponse struct {
	Values    []TestLibrary `json:"values"`
	Total     int           `json:"total"`
	PageSize  int           `json:"page_size"`
	PageIndex int           `json:"page_index"`
}

type TestLibraryMember struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type TestLibraryMemberListResponse struct {
	Values    []TestLibraryMember `json:"values"`
	Total     int                 `json:"total"`
	PageSize  int                 `json:"page_size"`
	PageIndex int                 `json:"page_index"`
}

type TestSuite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TestSuiteListResponse struct {
	Values    []TestSuite `json:"values"`
	Total     int         `json:"total"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
}

type TestCase struct {
	ID               string `json:"id"`
	Identifier       string `json:"identifier,omitempty"`
	Title            string `json:"title,omitempty"`
	LibraryID        string `json:"test_library_id,omitempty"`
	SuiteID          string `json:"suite_id,omitempty"`
	TypeID           string `json:"type_id,omitempty"`
	ImportantLevelID string `json:"important_level_id,omitempty"`
	MaintenanceID    string `json:"maintenance_id,omitempty"`
	Description      string `json:"description,omitempty"`
	Precondition     string `json:"precondition,omitempty"`
}

type TestCaseListResponse struct {
	Values    []TestCase `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type TestPlan struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TypeID     string `json:"type_id,omitempty"`
	AssigneeID string `json:"assignee_id,omitempty"`
	StartAt    int64  `json:"start_at,omitempty"`
	EndAt      int64  `json:"end_at,omitempty"`
	ProjectID  string `json:"project_id,omitempty"`
	SprintID   string `json:"sprint_id,omitempty"`
	VersionID  string `json:"version_id,omitempty"`
}

type TestPlanListResponse struct {
	Values    []TestPlan `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type TestRun struct {
	ID         string `json:"id"`
	LibraryID  string `json:"library_id"`
	PlanID     string `json:"plan_id"`
	CaseID     string `json:"case_id"`
	ExecutorID string `json:"executor_id,omitempty"`
}

// Library operations

func (c *Client) ListTestLibraries() ([]TestLibrary, error) {
	resp, err := c.Request("GET", "/testhub/libraries", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list test libraries failed with status: %d", resp.StatusCode)
	}
	var listResp TestLibraryListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) CreateTestLibrary(name, identifier, scopeType, scopeID, visibility, description string) (*TestLibrary, error) {
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
	resp, err := c.Request("POST", "/testhub/libraries", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create test library failed with status: %d", resp.StatusCode)
	}
	var lib TestLibrary
	if err := json.NewDecoder(resp.Body).Decode(&lib); err != nil {
		return nil, err
	}
	return &lib, nil
}

func (c *Client) UpdateTestLibrary(libraryID, name, visibility, description string) (*TestLibrary, error) {
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
	resp, err := c.Request("PATCH", fmt.Sprintf("/testhub/libraries/%s", libraryID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update test library failed with status: %d", resp.StatusCode)
	}
	var lib TestLibrary
	if err := json.NewDecoder(resp.Body).Decode(&lib); err != nil {
		return nil, err
	}
	return &lib, nil
}

// Library member operations

func (c *Client) ListTestLibraryMembers(libraryID string) ([]TestLibraryMember, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/testhub/libraries/%s/members", libraryID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list test library members failed with status: %d", resp.StatusCode)
	}
	var listResp TestLibraryMemberListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) AddTestLibraryMember(libraryID, memberID, memberType, roleID string) (*TestLibraryMember, error) {
	body := map[string]interface{}{
		"member": map[string]interface{}{
			"id":   memberID,
			"type": memberType,
		},
	}
	if roleID != "" {
		body["role_id"] = roleID
	}
	resp, err := c.Request("POST", fmt.Sprintf("/testhub/libraries/%s/members", libraryID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("add test library member failed with status: %d", resp.StatusCode)
	}
	var member TestLibraryMember
	if err := json.NewDecoder(resp.Body).Decode(&member); err != nil {
		return nil, err
	}
	return &member, nil
}

func (c *Client) RemoveTestLibraryMember(libraryID, memberID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/testhub/libraries/%s/members/%s", libraryID, memberID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("remove test library member failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Suite operations

func (c *Client) ListTestSuites(libraryID string) ([]TestSuite, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/testhub/libraries/%s/suites", libraryID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list test suites failed with status: %d", resp.StatusCode)
	}
	var listResp TestSuiteListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) CreateTestSuite(libraryID, name string) (*TestSuite, error) {
	body := map[string]interface{}{
		"name": name,
	}
	resp, err := c.Request("POST", fmt.Sprintf("/testhub/libraries/%s/suites", libraryID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create test suite failed with status: %d", resp.StatusCode)
	}
	var suite TestSuite
	if err := json.NewDecoder(resp.Body).Decode(&suite); err != nil {
		return nil, err
	}
	return &suite, nil
}

func (c *Client) DeleteTestSuite(libraryID, suiteID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/testhub/libraries/%s/suites/%s", libraryID, suiteID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete test suite failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Case operations

func (c *Client) ListTestCases(libraryID, suiteID string) ([]TestCase, error) {
	url := fmt.Sprintf("/testhub/cases?library_id=%s", libraryID)
	if suiteID != "" {
		url += "&suite_id=" + suiteID
	}
	resp, err := c.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list test cases failed with status: %d", resp.StatusCode)
	}
	var listResp TestCaseListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) GetTestCase(caseID string) (*TestCase, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/testhub/cases/%s", caseID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get test case failed with status: %d", resp.StatusCode)
	}
	var tc TestCase
	if err := json.NewDecoder(resp.Body).Decode(&tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (c *Client) CreateTestCase(tc *TestCase) (*TestCase, error) {
	body := map[string]interface{}{}
	if tc.LibraryID != "" {
		body["test_library_id"] = tc.LibraryID
	}
	if tc.Title != "" {
		body["title"] = tc.Title
	}
	if tc.SuiteID != "" {
		body["suite_id"] = tc.SuiteID
	}
	if tc.TypeID != "" {
		body["type_id"] = tc.TypeID
	}
	if tc.MaintenanceID != "" {
		body["maintenance_id"] = tc.MaintenanceID
	}
	if tc.Description != "" {
		body["description"] = tc.Description
	}
	if tc.Precondition != "" {
		body["precondition"] = tc.Precondition
	}
	resp, err := c.Request("POST", "/testhub/cases", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create test case failed with status: %d", resp.StatusCode)
	}
	var created TestCase
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

func (c *Client) UpdateTestCase(caseID string, tc *TestCase) (*TestCase, error) {
	body := map[string]interface{}{}
	if tc.Title != "" {
		body["title"] = tc.Title
	}
	if tc.SuiteID != "" {
		body["suite_id"] = tc.SuiteID
	}
	if tc.TypeID != "" {
		body["type_id"] = tc.TypeID
	}
	if tc.MaintenanceID != "" {
		body["maintenance_id"] = tc.MaintenanceID
	}
	if tc.Description != "" {
		body["description"] = tc.Description
	}
	if tc.Precondition != "" {
		body["precondition"] = tc.Precondition
	}
	resp, err := c.Request("PATCH", fmt.Sprintf("/testhub/cases/%s", caseID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update test case failed with status: %d", resp.StatusCode)
	}
	var updated TestCase
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (c *Client) DeleteTestCase(caseID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/testhub/cases/%s", caseID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete test case failed with status: %d", resp.StatusCode)
	}
	return nil
}

// Plan operations

func (c *Client) ListTestPlans(libraryID string) ([]TestPlan, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/testhub/libraries/%s/plans", libraryID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list test plans failed with status: %d", resp.StatusCode)
	}
	var listResp TestPlanListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp.Values, nil
}

func (c *Client) CreateTestPlan(libraryID string, plan *TestPlan) (*TestPlan, error) {
	body := map[string]interface{}{
		"name": plan.Name,
	}
	if plan.TypeID != "" {
		body["type_id"] = plan.TypeID
	}
	if plan.AssigneeID != "" {
		body["assignee_id"] = plan.AssigneeID
	}
	if plan.StartAt != 0 {
		body["start_at"] = plan.StartAt
	}
	if plan.EndAt != 0 {
		body["end_at"] = plan.EndAt
	}
	if plan.ProjectID != "" {
		body["project_id"] = plan.ProjectID
	}
	if plan.SprintID != "" {
		body["sprint_id"] = plan.SprintID
	}
	if plan.VersionID != "" {
		body["version_id"] = plan.VersionID
	}
	resp, err := c.Request("POST", fmt.Sprintf("/testhub/libraries/%s/plans", libraryID), body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create test plan failed with status: %d", resp.StatusCode)
	}
	var created TestPlan
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

// Run operations

func (c *Client) CreateTestRun(libraryID, planID, caseID, executorID string) (*TestRun, error) {
	body := map[string]interface{}{
		"library_id": libraryID,
		"plan_id":    planID,
		"case_id":    caseID,
	}
	if executorID != "" {
		body["executor_id"] = executorID
	}
	resp, err := c.Request("POST", "/testhub/runs", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create test run failed with status: %d", resp.StatusCode)
	}
	var run TestRun
	if err := json.NewDecoder(resp.Body).Decode(&run); err != nil {
		return nil, err
	}
	return &run, nil
}
