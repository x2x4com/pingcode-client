package project

import (
	"net/url"
	"pingcode-client/internal/pkg/sdk"
	"time"
)

// parseDateToMs converts "YYYY-MM-DD" to Unix milliseconds (UTC midnight).
// Returns 0 if the string is empty or unparseable.
func parseDateToMs(s string) int64 {
	if s == "" {
		return 0
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return 0
	}
	return t.UTC().Unix()
}

func ListProjects(client *sdk.Client) ([]sdk.Project, error) {
	return client.ListProjects()
}

func GetProject(client *sdk.Client, id string) (*sdk.Project, error) {
	return client.GetProject(id)
}

func CreateProject(client *sdk.Client, name, desc, pType string) (*sdk.Project, error) {
	p := &sdk.Project{Name: name, Desc: desc, Type: pType}
	return client.CreateProject(p)
}

func UpdateProject(client *sdk.Client, id, name, desc string) (*sdk.Project, error) {
	p := &sdk.Project{Name: name, Desc: desc}
	return client.UpdateProject(id, p)
}

func DeleteProject(client *sdk.Client, id string) error {
	return client.DeleteProject(id)
}

func GetWorkItem(client *sdk.Client, id string, identifier string) (*sdk.WorkItem, error) {
	return client.GetWorkItem(id, identifier)
}

func CreateWorkItem(client *sdk.Client, projectID, title, desc, iType, typeID, stateID, priorityID, assigneeID, sprintID, boardID, entryID, swimlaneID, parentID string, versionIDs []string) (*sdk.WorkItem, error) {
	item := &sdk.WorkItem{
		ProjectID:  projectID,
		Title:      title,
		Desc:       desc,
		TypeID:     typeID,
		StateID:    stateID,
		PriorityID: priorityID,
		AssigneeID: assigneeID,
		SprintID:   sprintID,
		BoardID:    boardID,
		EntryID:    entryID,
		SwimlaneID: swimlaneID,
		ParentID:   parentID,
		VersionIDs: versionIDs,
	}
	return client.CreateWorkItem(item)
}

func UpdateWorkItem(client *sdk.Client, id, title, desc, stateID, priorityID, assigneeID, sprintID, boardID, entryID, swimlaneID, parentID string, versionIDs []string) (*sdk.WorkItem, error) {
	item := &sdk.WorkItem{
		Title:      title,
		Desc:       desc,
		StateID:    stateID,
		PriorityID: priorityID,
		AssigneeID: assigneeID,
		SprintID:   sprintID,
		BoardID:    boardID,
		EntryID:    entryID,
		SwimlaneID: swimlaneID,
		ParentID:   parentID,
		VersionIDs: versionIDs,
	}
	return client.UpdateWorkItem(id, item)
}

func DeleteWorkItem(client *sdk.Client, id string) error {
	return client.DeleteWorkItem(id)
}

func AddWorkItemTag(client *sdk.Client, workItemID string, tagID string) error {
	return client.AddWorkItemTag(workItemID, tagID)
}

func RemoveWorkItemTag(client *sdk.Client, workItemID string, tagID string) error {
	return client.RemoveWorkItemTag(workItemID, tagID)
}

func GetIteration(client *sdk.Client, projectID, id string) (*sdk.Iteration, error) {
	return client.GetIteration(projectID, id)
}

func CreateIteration(client *sdk.Client, projectID, name, start, end, assigneeID, desc, status string) (*sdk.Iteration, error) {
	iter := &sdk.Iteration{
		ProjectID:   projectID,
		Name:        name,
		StartAt:     parseDateToMs(start),
		EndAt:       parseDateToMs(end),
		AssigneeID:  assigneeID,
		Description: desc,
		Status:      status,
	}
	return client.CreateIteration(iter)
}

func UpdateIteration(client *sdk.Client, projectID, id, name, start, end, assigneeID, desc, status string) (*sdk.Iteration, error) {
	iter := &sdk.Iteration{
		Name:        name,
		StartAt:     parseDateToMs(start),
		EndAt:       parseDateToMs(end),
		AssigneeID:  assigneeID,
		Description: desc,
		Status:      status,
	}
	return client.UpdateIteration(projectID, id, iter)
}

func DeleteIteration(client *sdk.Client, projectID, id string) error {
	return client.DeleteIteration(projectID, id)
}

func GetVersion(client *sdk.Client, projectID, id string) (*sdk.Version, error) {
	return client.GetVersion(projectID, id)
}

func CreateVersion(client *sdk.Client, projectID, name, start, end, assigneeID, stageID string) (*sdk.Version, error) {
	v := &sdk.Version{
		ProjectID:  projectID,
		Name:       name,
		StartAt:    parseDateToMs(start),
		EndAt:      parseDateToMs(end),
		AssigneeID: assigneeID,
		StageID:    stageID,
	}
	return client.CreateVersion(v)
}

func UpdateVersion(client *sdk.Client, projectID, id, name, start, end, assigneeID, stageID string) (*sdk.Version, error) {
	v := &sdk.Version{
		Name:       name,
		StartAt:    parseDateToMs(start),
		EndAt:      parseDateToMs(end),
		AssigneeID: assigneeID,
		StageID:    stageID,
	}
	return client.UpdateVersion(projectID, id, v)
}

func DeleteVersion(client *sdk.Client, projectID, id string) error {
	return client.DeleteVersion(projectID, id)
}

func ListKanbans(client *sdk.Client, projectID string) ([]sdk.Kanban, error) {
	return client.ListKanbans(projectID)
}

func ListKanbanEntries(client *sdk.Client, projectID string, boardID string) ([]sdk.KanbanEntry, error) {
	return client.ListKanbanEntries(projectID, boardID)
}

func ListProjectMembers(client *sdk.Client, projectID string) ([]sdk.ProjectMember, error) {
	return client.ListProjectMembers(projectID)
}

func GetKanban(client *sdk.Client, id string) (*sdk.Kanban, error) {
	return client.GetKanban(id)
}

func CreateKanban(client *sdk.Client, projectID, name string) (*sdk.Kanban, error) {
	k := &sdk.Kanban{ProjectID: projectID, Name: name}
	return client.CreateKanban(k)
}

func UpdateKanban(client *sdk.Client, id, name string) (*sdk.Kanban, error) {
	k := &sdk.Kanban{Name: name}
	return client.UpdateKanban(id, k)
}

func DeleteKanban(client *sdk.Client, id string) error {
	return client.DeleteKanban(id)
}

func ListWorkItems(client *sdk.Client, projectID string, tagIDs string, stateIDs string, sprintID string, assigneeID string) ([]sdk.WorkItem, error) {
	filters := url.Values{}
	if projectID != "" {
		filters.Add("project_ids", projectID)
	}
	if tagIDs != "" {
		filters.Add("tag_ids", tagIDs)
	}
	if stateIDs != "" {
		filters.Add("state_ids", stateIDs)
	}
	if sprintID != "" {
		filters.Add("sprint_id", sprintID)
	}
	if assigneeID != "" {
		filters.Add("assignee_id", assigneeID)
	}
	return client.ListWorkItems(filters)
}

func ListIterations(client *sdk.Client, projectID string) ([]sdk.Iteration, error) {
	return client.ListIterations(projectID)
}

func ListVersions(client *sdk.Client, projectID string) ([]sdk.Version, error) {
	return client.ListVersions(projectID)
}

func ListWorkItemTypes(client *sdk.Client) ([]sdk.WorkItemType, error) {
	return client.ListWorkItemTypes()
}

func ListPriorities(client *sdk.Client, projectID string) ([]sdk.Priority, error) {
	return client.ListPriorities(projectID)
}

func ListTags(client *sdk.Client) ([]sdk.Tag, error) {
	return client.ListTags()
}

func CreateWorkItemTag(client *sdk.Client, name string) (*sdk.Tag, error) {
	return client.CreateWorkItemTag(name)
}

func ListProperties(client *sdk.Client, projectID string, workItemTypeID string) ([]sdk.Property, error) {
	return client.ListProperties(projectID, workItemTypeID)
}

func ListStatuses(client *sdk.Client) ([]sdk.Status, error) {
	return client.ListStatuses()
}

func ListSeverities(client *sdk.Client) ([]sdk.Severity, error) {
	return client.ListSeverities()
}

// ── Swimlane ───────────────────────────────────────────────────────────────

func ListSwimlanes(client *sdk.Client, projectID, boardID string) ([]sdk.Swimlane, error) {
	return client.ListSwimlanes(projectID, boardID)
}

func CreateSwimlane(client *sdk.Client, projectID, boardID, name string) (*sdk.Swimlane, error) {
	return client.CreateSwimlane(projectID, boardID, name)
}

func UpdateSwimlane(client *sdk.Client, projectID, boardID, swimlaneID, name string) (*sdk.Swimlane, error) {
	return client.UpdateSwimlane(projectID, boardID, swimlaneID, name)
}

func DeleteSwimlane(client *sdk.Client, projectID, boardID, swimlaneID string) error {
	return client.DeleteSwimlane(projectID, boardID, swimlaneID)
}

// ── Kanban Entry CRUD ──────────────────────────────────────────────────────

func CreateKanbanEntry(client *sdk.Client, projectID, boardID, name string, wipLimit int, isSplit bool, dod string) (*sdk.KanbanEntry, error) {
	return client.CreateKanbanEntry(projectID, boardID, name, wipLimit, isSplit, dod)
}

func UpdateKanbanEntry(client *sdk.Client, projectID, boardID, entryID, name string, wipLimit int, isSplit bool, dod string) (*sdk.KanbanEntry, error) {
	return client.UpdateKanbanEntry(projectID, boardID, entryID, name, wipLimit, isSplit, dod)
}

func DeleteKanbanEntry(client *sdk.Client, projectID, boardID, entryID string) error {
	return client.DeleteKanbanEntry(projectID, boardID, entryID)
}

// ── Project Member CRUD ────────────────────────────────────────────────────

func AddProjectMember(client *sdk.Client, projectID, memberID, memberType, roleID string) (*sdk.ProjectMember, error) {
	return client.AddProjectMember(projectID, memberID, memberType, roleID)
}

func RemoveProjectMember(client *sdk.Client, projectID, memberID string) error {
	return client.RemoveProjectMember(projectID, memberID)
}

func UpdateProjectMember(client *sdk.Client, projectID, memberID, roleID string) (*sdk.ProjectMember, error) {
	return client.UpdateProjectMember(projectID, memberID, roleID)
}

// ── WorkItem Participant ───────────────────────────────────────────────────

func ListWorkItemParticipants(client *sdk.Client, workItemID string) ([]sdk.WorkItemParticipant, error) {
	return client.ListWorkItemParticipants(workItemID)
}

func AddWorkItemParticipant(client *sdk.Client, workItemID, userID string) (*sdk.WorkItemParticipant, error) {
	return client.AddWorkItemParticipant(workItemID, userID)
}

func RemoveWorkItemParticipant(client *sdk.Client, workItemID, participantID string) error {
	return client.RemoveWorkItemParticipant(workItemID, participantID)
}

// ── WorkItem Relation ──────────────────────────────────────────────────────

func ListWorkItemRelations(client *sdk.Client, workItemID string) ([]sdk.WorkItemRelation, error) {
	return client.ListWorkItemRelations(workItemID)
}

func AddWorkItemRelation(client *sdk.Client, workItemID, targetID, relationType string) (*sdk.WorkItemRelation, error) {
	return client.AddWorkItemRelation(workItemID, targetID, relationType)
}

func RemoveWorkItemRelation(client *sdk.Client, workItemID, relationID string) error {
	return client.RemoveWorkItemRelation(workItemID, relationID)
}

// ── WorkItem Transition History ────────────────────────────────────────────

func ListWorkItemTransitionHistories(client *sdk.Client, workItemID string) ([]sdk.WorkItemTransitionHistory, error) {
	return client.ListWorkItemTransitionHistories(workItemID)
}

// ── Sprint Sections ────────────────────────────────────────────────────────

func ListSprintSections(client *sdk.Client, projectID string) ([]sdk.SprintSection, error) {
	return client.ListSprintSections(projectID)
}

func CreateSprintSection(client *sdk.Client, projectID, name string) (*sdk.SprintSection, error) {
	return client.CreateSprintSection(projectID, name)
}

func UpdateSprintSection(client *sdk.Client, projectID, sectionID, name string) (*sdk.SprintSection, error) {
	return client.UpdateSprintSection(projectID, sectionID, name)
}

func DeleteSprintSection(client *sdk.Client, projectID, sectionID string) error {
	return client.DeleteSprintSection(projectID, sectionID)
}

func ListSprintCategories(client *sdk.Client, projectID string) ([]sdk.SprintCategory, error) {
	return client.ListSprintCategories(projectID)
}

func CreateSprintCategory(client *sdk.Client, projectID, name, sectionID string) (*sdk.SprintCategory, error) {
	return client.CreateSprintCategory(projectID, name, sectionID)
}

func UpdateSprintCategory(client *sdk.Client, projectID, categoryID, name string) (*sdk.SprintCategory, error) {
	return client.UpdateSprintCategory(projectID, categoryID, name)
}

func DeleteSprintCategory(client *sdk.Client, projectID, categoryID string) error {
	return client.DeleteSprintCategory(projectID, categoryID)
}

// ── Version Sections ───────────────────────────────────────────────────────

func ListVersionSections(client *sdk.Client, projectID string) ([]sdk.VersionSection, error) {
	return client.ListVersionSections(projectID)
}

func CreateVersionSection(client *sdk.Client, projectID, name string) (*sdk.VersionSection, error) {
	return client.CreateVersionSection(projectID, name)
}

func UpdateVersionSection(client *sdk.Client, projectID, sectionID, name string) (*sdk.VersionSection, error) {
	return client.UpdateVersionSection(projectID, sectionID, name)
}

func DeleteVersionSection(client *sdk.Client, projectID, sectionID string) error {
	return client.DeleteVersionSection(projectID, sectionID)
}

func ListVersionCategories(client *sdk.Client, projectID string) ([]sdk.VersionCategory, error) {
	return client.ListVersionCategories(projectID)
}

func CreateVersionCategory(client *sdk.Client, projectID, name, sectionID string) (*sdk.VersionCategory, error) {
	return client.CreateVersionCategory(projectID, name, sectionID)
}

func UpdateVersionCategory(client *sdk.Client, projectID, categoryID, name string) (*sdk.VersionCategory, error) {
	return client.UpdateVersionCategory(projectID, categoryID, name)
}

func DeleteVersionCategory(client *sdk.Client, projectID, categoryID string) error {
	return client.DeleteVersionCategory(projectID, categoryID)
}
