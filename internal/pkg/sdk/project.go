package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type,omitempty"` // scrum or kanban
}

type Iteration struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProjectID   string `json:"project_id,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	AssigneeID  string `json:"assignee_id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type Version struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProjectID  string `json:"project_id,omitempty"`
	StartAt    int64  `json:"start_at,omitempty"`
	EndAt      int64  `json:"end_at,omitempty"`
	AssigneeID string `json:"assignee_id,omitempty"`
	StageID    string `json:"stage_id,omitempty"`
}

type Kanban struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ProjectID string `json:"project_id,omitempty"`
}

type KanbanEntry struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	WipLimit         int    `json:"wip_limit,omitempty"`
	IsSplit          bool   `json:"is_split,omitempty"`
	DefinitionOfDone string `json:"definition_of_done,omitempty"`
}

type KanbanEntryListResponse struct {
	Values    []KanbanEntry `json:"values"`
	Total     int           `json:"total"`
	PageSize  int           `json:"page_size"`
	PageIndex int           `json:"page_index"`
}

type Swimlane struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SwimlaneListResponse struct {
	Values    []Swimlane `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type ProjectMember struct {
	ID   string   `json:"id"`
	User *UserRef `json:"user,omitempty"`
}

type ProjectMemberListResponse struct {
	Values    []ProjectMember `json:"values"`
	Total     int             `json:"total"`
	PageSize  int             `json:"page_size"`
	PageIndex int             `json:"page_index"`
}

type UserRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WorkItem struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Title      string `json:"title,omitempty"`
	Desc       string `json:"description,omitempty"`
	ProjectID  string `json:"project_id,omitempty"`
	TypeID     string `json:"type_id,omitempty"`
	Type       string `json:"type,omitempty"`   // For response
	Status     string `json:"status,omitempty"` // For response (sometimes string, sometimes object)

	// Creation & Extended fields
	StartAt           int64          `json:"start_at,omitempty"`
	EndAt             int64          `json:"end_at,omitempty"`
	PriorityID        string         `json:"priority_id,omitempty"`
	StateID           string         `json:"state_id,omitempty"`
	AssigneeID        string         `json:"assignee_id,omitempty"`
	ParentID          string         `json:"parent_id,omitempty"`
	SprintID          string         `json:"sprint_id,omitempty"`
	VersionIDs        []string       `json:"version_ids,omitempty"`
	BoardID           string         `json:"board_id,omitempty"`
	EntryID           string         `json:"entry_id,omitempty"`
	SwimlaneID        string         `json:"swimlane_id,omitempty"`
	StoryPoints       float64        `json:"story_points,omitempty"`
	EstimatedWorkload float64        `json:"estimated_workload,omitempty"`
	RemainingWorkload float64        `json:"remaining_workload,omitempty"`
	Properties        map[string]any `json:"properties,omitempty"`
	ParticipantIDs    []string       `json:"participant_ids,omitempty"`
	Tags              []Tag          `json:"tags,omitempty"`

	// Full objects for response
	Assignee    *UserRef     `json:"assignee,omitempty"`
	PriorityObj *Priority    `json:"priority,omitempty"`
	StateObj    *Status      `json:"state,omitempty"`
	SprintObj   *Iteration   `json:"sprint,omitempty"`
	BoardObj    *Kanban      `json:"board,omitempty"`
	EntryObj    *KanbanEntry `json:"entry,omitempty"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Property struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Type    string           `json:"type"`
	Options []PropertyOption `json:"options,omitempty"`
}

type PropertyOption struct {
	ID   string `json:"_id"`
	Text string `json:"text"`
}

type WorkItemType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Priority struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Severity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectListResponse struct {
	Values    []Project `json:"values"`
	Total     int       `json:"total"`
	PageSize  int       `json:"page_size"`
	PageIndex int       `json:"page_index"`
}

type IterationListResponse struct {
	Values []Iteration `json:"values"`
	Total  int         `json:"total"`
}

type VersionListResponse struct {
	Values []Version `json:"values"`
	Total  int       `json:"total"`
}

type KanbanListResponse struct {
	Values []Kanban `json:"values"`
	Total  int      `json:"total"`
}

type WorkItemListResponse struct {
	Values []WorkItem `json:"values"`
	Total  int        `json:"total"`
}

type WorkItemTypeListResponse struct {
	Values []WorkItemType `json:"values"`
}

type PriorityListResponse struct {
	Values    []Priority `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type TagListResponse struct {
	Values    []Tag `json:"values"`
	Total     int   `json:"total"`
	PageSize  int   `json:"page_size"`
	PageIndex int   `json:"page_index"`
}

type PropertyListResponse struct {
	Values    []Property `json:"values"`
	Total     int        `json:"total"`
	PageSize  int        `json:"page_size"`
	PageIndex int        `json:"page_index"`
}

type StatusListResponse struct {
	Values []Status `json:"values"`
}

type SeverityListResponse struct {
	Values []Severity `json:"values"`
}

func (c *Client) ListProjects() ([]Project, error) {
	resp, err := c.Request("GET", "/project/projects", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list projects failed with status: %d", resp.StatusCode)
	}

	var listResp ProjectListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) GetProject(id string) (*Project, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get project failed with status: %d", resp.StatusCode)
	}

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) CreateProject(project *Project) (*Project, error) {
	resp, err := c.Request("POST", "/project/projects", project)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create project failed with status: %d", resp.StatusCode)
	}

	var created Project
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (c *Client) UpdateProject(id string, project *Project) (*Project, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/project/projects/%s", id), project)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update project failed with status: %d", resp.StatusCode)
	}

	var updated Project
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *Client) DeleteProject(id string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/projects/%s", id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete project failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListStatuses() ([]Status, error) {
	resp, err := c.Request("GET", "/project/work_item_states", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list statuses failed with status: %d", resp.StatusCode)
	}

	var listResp StatusListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListSeverities() ([]Severity, error) {
	resp, err := c.Request("GET", "/project/severities", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list severities failed with status: %d", resp.StatusCode)
	}

	var listResp SeverityListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListWorkItemTypes() ([]WorkItemType, error) {
	resp, err := c.Request("GET", "/project/work_item_types", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list workitem types failed with status: %d", resp.StatusCode)
	}

	var listResp WorkItemTypeListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListPriorities(projectID string) ([]Priority, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project_id is required for listing priorities")
	}
	resp, err := c.Request("GET", fmt.Sprintf("/project/work_item/priorities?project_id=%s", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list priorities failed with status: %d", resp.StatusCode)
	}

	var listResp PriorityListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListTags() ([]Tag, error) {
	resp, err := c.Request("GET", "/project/work_item_tags", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list tags failed with status: %d", resp.StatusCode)
	}

	var listResp TagListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) CreateWorkItemTag(name string) (*Tag, error) {
	if name == "" {
		return nil, fmt.Errorf("tag name is required")
	}
	payload := map[string]string{"name": name}
	resp, err := c.Request("POST", "/project/work_item_tags", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("create tag failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var tag Tag
	if err := json.NewDecoder(resp.Body).Decode(&tag); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (c *Client) ListProperties(projectID string, workItemTypeID string) ([]Property, error) {
	if projectID == "" || workItemTypeID == "" {
		return nil, fmt.Errorf("project_id and work_item_type_id are required for listing properties")
	}
	resp, err := c.Request("GET", fmt.Sprintf("/project/work_item/properties?project_id=%s&work_item_type_id=%s", projectID, workItemTypeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list properties failed with status: %d", resp.StatusCode)
	}

	var listResp PropertyListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) AddWorkItemTag(workItemID string, tagID string) error {
	if workItemID == "" || tagID == "" {
		return fmt.Errorf("work_item_id and tag_id are required")
	}
	payload := map[string]string{"tag_id": tagID}
	resp, err := c.Request("POST", fmt.Sprintf("/project/work_items/%s/tags", workItemID), payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("add tag failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (c *Client) RemoveWorkItemTag(workItemID string, tagID string) error {
	if workItemID == "" || tagID == "" {
		return fmt.Errorf("work_item_id and tag_id are required")
	}
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/work_items/%s/tags/%s", workItemID, tagID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("remove tag failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (c *Client) ListWorkItems(filters url.Values) ([]WorkItem, error) {
	path := "/project/work_items"
	if len(filters) > 0 {
		path = fmt.Sprintf("%s?%s", path, filters.Encode())
	}
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list workitems failed with status: %d", resp.StatusCode)
	}

	var listResp WorkItemListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) CreateWorkItem(item *WorkItem) (*WorkItem, error) {
	resp, err := c.Request("POST", "/project/work_items", item)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("create workitem failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var created WorkItem
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (c *Client) GetWorkItem(id string, identifier string) (*WorkItem, error) {
	if id != "" && identifier != "" {
		return nil, fmt.Errorf("cannot specify both id and identifier, use one only")
	}

	if id != "" {
		resp, err := c.Request("GET", fmt.Sprintf("/project/work_items/%s", id), nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("get workitem failed with status: %d", resp.StatusCode)
		}
		var item WorkItem
		if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	if identifier != "" {
		// Fallback: search by identifier
		items, err := c.ListWorkItems(nil)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if item.Identifier == identifier {
				return &item, nil
			}
		}
		return nil, fmt.Errorf("workitem not found with identifier: %s", identifier)
	}

	return nil, fmt.Errorf("workitem id or identifier is required")
}

func (c *Client) UpdateWorkItem(id string, item *WorkItem) (*WorkItem, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/project/work_items/%s", id), item)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("update workitem failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var updated WorkItem
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *Client) DeleteWorkItem(id string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/work_items/%s", id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete workitem failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListIterations(projectID string) ([]Iteration, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/sprints", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list iterations failed with status: %d", resp.StatusCode)
	}

	var listResp IterationListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) GetIteration(id string) (*Iteration, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/sprints/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get iteration failed with status: %d", resp.StatusCode)
	}

	var iteration Iteration
	if err := json.NewDecoder(resp.Body).Decode(&iteration); err != nil {
		return nil, err
	}

	return &iteration, nil
}

func (c *Client) CreateIteration(iteration *Iteration) (*Iteration, error) {
	resp, err := c.Request("POST", "/project/sprints", iteration)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create iteration failed with status: %d", resp.StatusCode)
	}

	var created Iteration
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (c *Client) UpdateIteration(id string, iteration *Iteration) (*Iteration, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/project/sprints/%s", id), iteration)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update iteration failed with status: %d", resp.StatusCode)
	}

	var updated Iteration
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *Client) DeleteIteration(id string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/sprints/%s", id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete iteration failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListVersions(projectID string) ([]Version, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/versions", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list versions failed with status: %d", resp.StatusCode)
	}

	var listResp VersionListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) GetVersion(id string) (*Version, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/versions/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get version failed with status: %d", resp.StatusCode)
	}

	var version Version
	if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
		return nil, err
	}

	return &version, nil
}

func (c *Client) CreateVersion(version *Version) (*Version, error) {
	resp, err := c.Request("POST", "/project/versions", version)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create version failed with status: %d", resp.StatusCode)
	}

	var created Version
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (c *Client) UpdateVersion(id string, version *Version) (*Version, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/project/versions/%s", id), version)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update version failed with status: %d", resp.StatusCode)
	}

	var updated Version
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *Client) DeleteVersion(id string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/versions/%s", id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete version failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListKanbans(projectID string) ([]Kanban, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/boards", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list kanbans failed with status: %d", resp.StatusCode)
	}

	var listResp KanbanListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListKanbanEntries(projectID string, boardID string) ([]KanbanEntry, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/boards/%s/entries", projectID, boardID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list kanban entries failed with status: %d", resp.StatusCode)
	}

	var listResp KanbanEntryListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListProjectMembers(projectID string) ([]ProjectMember, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project_id is required")
	}
	resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/members", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list project members failed with status: %d", resp.StatusCode)
	}

	var listResp ProjectMemberListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) GetKanban(id string) (*Kanban, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/project/boards/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get kanban failed with status: %d", resp.StatusCode)
	}

	var kanban Kanban
	if err := json.NewDecoder(resp.Body).Decode(&kanban); err != nil {
		return nil, err
	}

	return &kanban, nil
}

func (c *Client) CreateKanban(kanban *Kanban) (*Kanban, error) {
	resp, err := c.Request("POST", "/project/boards", kanban)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create kanban failed with status: %d", resp.StatusCode)
	}

	var created Kanban
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (c *Client) UpdateKanban(id string, kanban *Kanban) (*Kanban, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/project/boards/%s", id), kanban)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update kanban failed with status: %d", resp.StatusCode)
	}

	var updated Kanban
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *Client) DeleteKanban(id string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/project/boards/%s", id), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete kanban failed with status: %d", resp.StatusCode)
	}

	return nil
}

// ── Swimlane CRUD ──────────────────────────────────────────────────────────

func (c *Client) ListSwimlanes(projectID, boardID string) ([]Swimlane, error) {
resp, err := c.Request("GET", fmt.Sprintf("/project/projects/%s/boards/%s/swimlanes", projectID, boardID), nil)
if err != nil {
return nil, err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
return nil, fmt.Errorf("list swimlanes failed with status: %d", resp.StatusCode)
}
var listResp SwimlaneListResponse
if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
return nil, err
}
return listResp.Values, nil
}

func (c *Client) CreateSwimlane(projectID, boardID, name string) (*Swimlane, error) {
body := map[string]interface{}{"name": name}
resp, err := c.Request("POST", fmt.Sprintf("/project/projects/%s/boards/%s/swimlanes", projectID, boardID), body)
if err != nil {
return nil, err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
return nil, fmt.Errorf("create swimlane failed with status: %d", resp.StatusCode)
}
var s Swimlane
if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
return nil, err
}
return &s, nil
}

func (c *Client) UpdateSwimlane(projectID, boardID, swimlaneID, name string) (*Swimlane, error) {
body := map[string]interface{}{}
if name != "" {
body["name"] = name
}
resp, err := c.Request("PATCH", fmt.Sprintf("/project/projects/%s/boards/%s/swimlanes/%s", projectID, boardID, swimlaneID), body)
if err != nil {
return nil, err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
return nil, fmt.Errorf("update swimlane failed with status: %d", resp.StatusCode)
}
var s Swimlane
if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
return nil, err
}
return &s, nil
}

func (c *Client) DeleteSwimlane(projectID, boardID, swimlaneID string) error {
resp, err := c.Request("DELETE", fmt.Sprintf("/project/projects/%s/boards/%s/swimlanes/%s", projectID, boardID, swimlaneID), nil)
if err != nil {
return err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
return fmt.Errorf("delete swimlane failed with status: %d", resp.StatusCode)
}
return nil
}

// ── Kanban Entry CRUD ──────────────────────────────────────────────────────

func (c *Client) CreateKanbanEntry(projectID, boardID, name string, wipLimit int, isSplit bool, dod string) (*KanbanEntry, error) {
body := map[string]interface{}{"name": name}
if wipLimit > 0 {
body["wip_limit"] = wipLimit
}
if isSplit {
body["is_split"] = true
}
if dod != "" {
body["definition_of_done"] = dod
}
resp, err := c.Request("POST", fmt.Sprintf("/project/projects/%s/boards/%s/entries", projectID, boardID), body)
if err != nil {
return nil, err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
return nil, fmt.Errorf("create kanban entry failed with status: %d", resp.StatusCode)
}
var e KanbanEntry
if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
return nil, err
}
return &e, nil
}

func (c *Client) UpdateKanbanEntry(projectID, boardID, entryID, name string, wipLimit int, isSplit bool, dod string) (*KanbanEntry, error) {
body := map[string]interface{}{}
if name != "" {
body["name"] = name
}
if wipLimit > 0 {
body["wip_limit"] = wipLimit
}
if isSplit {
body["is_split"] = true
}
if dod != "" {
body["definition_of_done"] = dod
}
resp, err := c.Request("PATCH", fmt.Sprintf("/project/projects/%s/boards/%s/entries/%s", projectID, boardID, entryID), body)
if err != nil {
return nil, err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
return nil, fmt.Errorf("update kanban entry failed with status: %d", resp.StatusCode)
}
var e KanbanEntry
if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
return nil, err
}
return &e, nil
}

func (c *Client) DeleteKanbanEntry(projectID, boardID, entryID string) error {
resp, err := c.Request("DELETE", fmt.Sprintf("/project/projects/%s/boards/%s/entries/%s", projectID, boardID, entryID), nil)
if err != nil {
return err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
return fmt.Errorf("delete kanban entry failed with status: %d", resp.StatusCode)
}
return nil
}
