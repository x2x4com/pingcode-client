package project

import (
	"fmt"
	"log"
	"net/url"
	"pingcode-client/internal/pkg/sdk"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// parseDateToMs converts "YYYY-MM-DD" to Unix milliseconds (UTC midnight).
// Returns 0 if the string is empty or unparseable.
func parseDateToMs(s string) int64 {
	if s == "" {
		return 0
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Fatalf("Invalid date format %q, expected YYYY-MM-DD", s)
	}
	return t.UTC().Unix()
}

func ListProjects(client *sdk.Client) {
	log.Println("Listing projects...")
	projects, err := client.ListProjects()
	if err != nil {
		log.Fatalf("Error listing projects: %v", err)
	}

	fmt.Println("Projects:")
	for _, p := range projects {
		fmt.Printf("- [%s] %s (%s)\n", p.ID, p.Name, p.Type)
	}
}

func GetProject(client *sdk.Client, id string) {
	if id == "" {
		log.Fatal("Project ID is required")
	}
	p, err := client.GetProject(id)
	if err != nil {
		log.Fatalf("Error getting project: %v", err)
	}
	fmt.Printf("Project: %s\nID: %s\nType: %s\nDesc: %s\n", p.Name, p.ID, p.Type, p.Desc)
}

func CreateProject(client *sdk.Client, name, desc, pType string) {
	if name == "" {
		log.Fatal("Name is required")
	}
	p := &sdk.Project{Name: name, Desc: desc, Type: pType}
	created, err := client.CreateProject(p)
	if err != nil {
		log.Fatalf("Error creating project: %v", err)
	}
	fmt.Printf("Created Project: %s (ID: %s)\n", created.Name, created.ID)
}

func UpdateProject(client *sdk.Client, id, name, desc string) {
	if id == "" {
		log.Fatal("Project ID is required")
	}
	p := &sdk.Project{Name: name, Desc: desc}
	updated, err := client.UpdateProject(id, p)
	if err != nil {
		log.Fatalf("Error updating project: %v", err)
	}
	fmt.Printf("Updated Project: %s (ID: %s)\n", updated.Name, updated.ID)
}

func DeleteProject(client *sdk.Client, id string) {
	if id == "" {
		log.Fatal("Project ID is required")
	}
	err := client.DeleteProject(id)
	if err != nil {
		log.Fatalf("Error deleting project: %v", err)
	}
	fmt.Printf("Deleted Project: %s\n", id)
}

func GetWorkItem(client *sdk.Client, id string, identifier string) {
	if id == "" && identifier == "" {
		log.Fatal("WorkItem ID or Identifier is required")
	}
	item, err := client.GetWorkItem(id, identifier)
	if err != nil {
		log.Fatalf("Error getting work item: %v", err)
	}
	var priorityName, stateName, assigneeName, boardName, entryName string
	if item.PriorityObj != nil {
		priorityName = item.PriorityObj.Name
	}
	if item.StateObj != nil {
		stateName = item.StateObj.Name
	}
	if item.Assignee != nil {
		assigneeName = item.Assignee.Name
	}
	if item.BoardObj != nil {
		boardName = item.BoardObj.Name
	}
	if item.EntryObj != nil {
		entryName = item.EntryObj.Name
	}

	var tagNames []string
	for _, t := range item.Tags {
		tagNames = append(tagNames, t.Name)
	}
	tagsStr := "无"
	if len(tagNames) > 0 {
		tagsStr = strings.Join(tagNames, ", ")
	}

	fmt.Printf("WorkItem: %s\nID: %s\nIdentifier: %s\nType: %s\nStatus: %s\nPriority: %s\nAssignee: %s\nBoard: %s\nEntry: %s\nTags: %s\n",
		item.Title, item.ID, item.Identifier, item.Type, stateName, priorityName, assigneeName, boardName, entryName, tagsStr)
}

func CreateWorkItem(client *sdk.Client, projectID, title, desc, iType, typeID, stateID, priorityID, assigneeID, sprintID, boardID, entryID, swimlaneID, parentID string, versionIDs []string) {
	if projectID == "" || title == "" {
		log.Fatal("Project ID and Title are required")
	}
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
	created, err := client.CreateWorkItem(item)
	if err != nil {
		log.Fatalf("Error creating work item: %v", err)
	}
	fmt.Printf("Created WorkItem: %s (ID: %s, Identifier: %s)\n", created.Title, created.ID, created.Identifier)
}

func UpdateWorkItem(client *sdk.Client, id, title, desc, stateID, priorityID, assigneeID, sprintID, boardID, entryID, swimlaneID, parentID string, versionIDs []string) {
	if id == "" {
		log.Fatal("WorkItem ID is required")
	}
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
	updated, err := client.UpdateWorkItem(id, item)
	if err != nil {
		log.Fatalf("Error updating work item: %v", err)
	}
	fmt.Printf("Updated WorkItem: %s (ID: %s)\n", updated.Title, updated.ID)
}

func DeleteWorkItem(client *sdk.Client, id string) {
	if id == "" {
		log.Fatal("WorkItem ID is required")
	}
	err := client.DeleteWorkItem(id)
	if err != nil {
		log.Fatalf("Error deleting work item: %v", err)
	}
	fmt.Printf("Deleted WorkItem: %s\n", id)
}

func AddWorkItemTag(client *sdk.Client, workItemID string, tagID string) {
	if workItemID == "" || tagID == "" {
		log.Fatal("WorkItem ID and Tag ID are required")
	}
	err := client.AddWorkItemTag(workItemID, tagID)
	if err != nil {
		log.Fatalf("Error adding tag to work item: %v", err)
	}
	fmt.Printf("Added Tag %s to WorkItem %s\n", tagID, workItemID)
}

func RemoveWorkItemTag(client *sdk.Client, workItemID string, tagID string) {
	if workItemID == "" || tagID == "" {
		log.Fatal("WorkItem ID and Tag ID are required")
	}
	err := client.RemoveWorkItemTag(workItemID, tagID)
	if err != nil {
		log.Fatalf("Error removing tag from work item: %v", err)
	}
	fmt.Printf("Removed Tag %s from WorkItem %s\n", tagID, workItemID)
}

func GetIteration(client *sdk.Client, projectID, id string) {
	if id == "" {
		log.Fatal("Iteration ID is required")
	}
	iter, err := client.GetIteration(projectID, id)
	if err != nil {
		log.Fatalf("Error getting iteration: %v", err)
	}
	fmt.Printf("Iteration: %s\nID: %s\nStart: %d\nEnd: %d\n", iter.Name, iter.ID, iter.StartAt, iter.EndAt)
}

func CreateIteration(client *sdk.Client, projectID, name, start, end, assigneeID, desc, status string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and Name are required")
	}
	iter := &sdk.Iteration{
		ProjectID:   projectID,
		Name:        name,
		StartAt:     parseDateToMs(start),
		EndAt:       parseDateToMs(end),
		AssigneeID:  assigneeID,
		Description: desc,
		Status:      status,
	}
	created, err := client.CreateIteration(iter)
	if err != nil {
		log.Fatalf("Error creating iteration: %v", err)
	}
	fmt.Printf("Created Iteration: %s (ID: %s)\n", created.Name, created.ID)
}

func UpdateIteration(client *sdk.Client, projectID, id, name, start, end, assigneeID, desc, status string) {
	if id == "" {
		log.Fatal("Iteration ID is required")
	}
	iter := &sdk.Iteration{
		Name:        name,
		StartAt:     parseDateToMs(start),
		EndAt:       parseDateToMs(end),
		AssigneeID:  assigneeID,
		Description: desc,
		Status:      status,
	}
	updated, err := client.UpdateIteration(projectID, id, iter)
	if err != nil {
		log.Fatalf("Error updating iteration: %v", err)
	}
	fmt.Printf("Updated Iteration: %s (ID: %s)\n", updated.Name, updated.ID)
}

func DeleteIteration(client *sdk.Client, projectID, id string) {
	if id == "" {
		log.Fatal("Iteration ID is required")
	}
	err := client.DeleteIteration(projectID, id)
	if err != nil {
		log.Fatalf("Error deleting iteration: %v", err)
	}
	fmt.Printf("Deleted Iteration: %s\n", id)
}

func GetVersion(client *sdk.Client, projectID, id string) {
	if id == "" {
		log.Fatal("Version ID is required")
	}
	v, err := client.GetVersion(projectID, id)
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}
	fmt.Printf("Version: %s\nID: %s\nEndAt: %d\n", v.Name, v.ID, v.EndAt)
}

func CreateVersion(client *sdk.Client, projectID, name, start, end, assigneeID, stageID string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and Name are required")
	}
	v := &sdk.Version{
		ProjectID:  projectID,
		Name:       name,
		StartAt:    parseDateToMs(start),
		EndAt:      parseDateToMs(end),
		AssigneeID: assigneeID,
		StageID:    stageID,
	}
	created, err := client.CreateVersion(v)
	if err != nil {
		log.Fatalf("Error creating version: %v", err)
	}
	fmt.Printf("Created Version: %s (ID: %s)\n", created.Name, created.ID)
}

func UpdateVersion(client *sdk.Client, projectID, id, name, start, end, assigneeID, stageID string) {
	if id == "" {
		log.Fatal("Version ID is required")
	}
	v := &sdk.Version{
		Name:       name,
		StartAt:    parseDateToMs(start),
		EndAt:      parseDateToMs(end),
		AssigneeID: assigneeID,
		StageID:    stageID,
	}
	updated, err := client.UpdateVersion(projectID, id, v)
	if err != nil {
		log.Fatalf("Error updating version: %v", err)
	}
	fmt.Printf("Updated Version: %s (ID: %s)\n", updated.Name, updated.ID)
}

func DeleteVersion(client *sdk.Client, projectID, id string) {
	if id == "" {
		log.Fatal("Version ID is required")
	}
	err := client.DeleteVersion(projectID, id)
	if err != nil {
		log.Fatalf("Error deleting version: %v", err)
	}
	fmt.Printf("Deleted Version: %s\n", id)
}

func ListKanbans(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required for listing kanbans")
	}
	log.Printf("Listing kanbans for project %s...", projectID)
	kanbans, err := client.ListKanbans(projectID)
	if err != nil {
		log.Fatalf("Error listing kanbans: %v", err)
	}

	fmt.Println("Kanbans:")
	for _, k := range kanbans {
		fmt.Printf("- %s (ID: %s)\n", k.Name, k.ID)
	}
}

func ListKanbanEntries(client *sdk.Client, projectID string, boardID string) {
	if projectID == "" || boardID == "" {
		log.Fatal("Project ID and Board ID are required")
	}
	log.Printf("Listing kanban entries for project %s, board %s...\n", projectID, boardID)
	entries, err := client.ListKanbanEntries(projectID, boardID)
	if err != nil {
		log.Fatalf("Error listing kanban entries: %v", err)
	}

	fmt.Println("Kanban Entries:")
	for _, e := range entries {
		fmt.Printf("- %s (ID: %s)\n", e.Name, e.ID)
	}
}

func ListProjectMembers(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	log.Printf("Listing members for project %s...\n", projectID)
	members, err := client.ListProjectMembers(projectID)
	if err != nil {
		log.Fatalf("Error listing project members: %v", err)
	}

	fmt.Println("Project Members:")
	for _, m := range members {
		if m.User != nil {
			fmt.Printf("- %s (ID: %s)\n", m.User.Name, m.User.ID)
		}
	}
}

func GetKanban(client *sdk.Client, id string) {
	if id == "" {
		log.Fatal("Kanban ID is required")
	}
	k, err := client.GetKanban(id)
	if err != nil {
		log.Fatalf("Error getting kanban: %v", err)
	}
	fmt.Printf("Kanban: %s\nID: %s\n", k.Name, k.ID)
}

func CreateKanban(client *sdk.Client, projectID, name string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and Name are required")
	}
	k := &sdk.Kanban{ProjectID: projectID, Name: name}
	created, err := client.CreateKanban(k)
	if err != nil {
		log.Fatalf("Error creating kanban: %v", err)
	}
	fmt.Printf("Created Kanban: %s (ID: %s)\n", created.Name, created.ID)
}

func UpdateKanban(client *sdk.Client, id, name string) {
	if id == "" {
		log.Fatal("Kanban ID is required")
	}
	k := &sdk.Kanban{Name: name}
	updated, err := client.UpdateKanban(id, k)
	if err != nil {
		log.Fatalf("Error updating kanban: %v", err)
	}
	fmt.Printf("Updated Kanban: %s (ID: %s)\n", updated.Name, updated.ID)
}

func DeleteKanban(client *sdk.Client, id string) {
	if id == "" {
		log.Fatal("Kanban ID is required")
	}
	err := client.DeleteKanban(id)
	if err != nil {
		log.Fatalf("Error deleting kanban: %v", err)
	}
	fmt.Printf("Deleted Kanban: %s\n", id)
}

func ListWorkItems(client *sdk.Client, projectID string, tagIDs string, stateIDs string) {
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

	log.Println("Listing work items...")
	items, err := client.ListWorkItems(filters)
	if err != nil {
		log.Fatalf("Error listing work items: %v", err)
	}

	fmt.Println("WorkItems:")
	for _, item := range items {
		fmt.Printf("- [%s] %s (Identifier: %s, Type: %s)\n", item.ID, item.Title, item.Identifier, item.Type)
	}
}

func ListIterations(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required for listing iterations")
	}
	log.Printf("Listing iterations for project %s...", projectID)
	iterations, err := client.ListIterations(projectID)
	if err != nil {
		log.Fatalf("Error listing iterations: %v", err)
	}

	fmt.Println("Iterations:")
	for _, iter := range iterations {
		fmt.Printf("- %s (ID: %s)\n", iter.Name, iter.ID)
	}
}

func ListVersions(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required for listing versions")
	}
	log.Printf("Listing versions for project %s...", projectID)
	versions, err := client.ListVersions(projectID)
	if err != nil {
		log.Fatalf("Error listing versions: %v", err)
	}

	fmt.Println("Versions:")
	for _, v := range versions {
		fmt.Printf("- %s (ID: %s)\n", v.Name, v.ID)
	}
}

func ListWorkItemTypes(client *sdk.Client) {
	log.Println("Listing work item types...")
	types, err := client.ListWorkItemTypes()
	if err != nil {
		log.Fatalf("Error listing work item types: %v", err)
	}

	fmt.Println("Work Item Types:")
	for _, t := range types {
		fmt.Printf("- %s (ID: %s)\n", t.Name, t.ID)
	}
}

func ListPriorities(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	log.Println("Listing priorities...")
	priorities, err := client.ListPriorities(projectID)
	if err != nil {
		log.Fatalf("Error listing priorities: %v", err)
	}

	fmt.Println("Priorities:")
	for _, p := range priorities {
		fmt.Printf("- %s (ID: %s)\n", p.Name, p.ID)
	}
}

func ListTags(client *sdk.Client) {
	log.Println("Listing tags...")
	tags, err := client.ListTags()
	if err != nil {
		log.Fatalf("Error listing tags: %v", err)
	}

	fmt.Println("Tags:")
	for _, t := range tags {
		fmt.Printf("- %s (ID: %s)\n", t.Name, t.ID)
	}
}

func CreateWorkItemTag(client *sdk.Client, name string) {
	if name == "" {
		log.Fatal("Tag name is required")
	}
	tag, err := client.CreateWorkItemTag(name)
	if err != nil {
		log.Fatalf("Error creating tag: %v", err)
	}
	fmt.Printf("Created Tag: %s (ID: %s)\n", tag.Name, tag.ID)
}

func ListProperties(client *sdk.Client, projectID string, workItemTypeID string) {
	if projectID == "" || workItemTypeID == "" {
		log.Fatal("Project ID and WorkItem Type ID are required")
	}
	log.Println("Listing properties...")
	properties, err := client.ListProperties(projectID, workItemTypeID)
	if err != nil {
		log.Fatalf("Error listing properties: %v", err)
	}

	fmt.Println("Properties:")
	for _, p := range properties {
		fmt.Printf("- %s (ID: %s, Type: %s)\n", p.Name, p.ID, p.Type)
		if len(p.Options) > 0 {
			fmt.Println("  Options:")
			for _, opt := range p.Options {
				fmt.Printf("    - %s (ID: %s)\n", opt.Text, opt.ID)
			}
		}
	}
}

func ListStatuses(client *sdk.Client) {
	log.Println("Listing statuses...")
	statuses, err := client.ListStatuses()
	if err != nil {
		log.Fatalf("Error listing statuses: %v", err)
	}

	fmt.Println("Statuses:")
	for _, s := range statuses {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func ListSeverities(client *sdk.Client) {
	log.Println("Listing severities...")
	severities, err := client.ListSeverities()
	if err != nil {
		log.Fatalf("Error listing severities: %v", err)
	}

	fmt.Println("Severities:")
	for _, s := range severities {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

// ── Swimlane ───────────────────────────────────────────────────────────────

func ListSwimlanes(client *sdk.Client, projectID, boardID string) {
	swimlanes, err := client.ListSwimlanes(projectID, boardID)
	if err != nil {
		log.Fatalf("Error listing swimlanes: %v", err)
	}
	fmt.Println("Swimlanes:")
	for _, s := range swimlanes {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func CreateSwimlane(client *sdk.Client, projectID, boardID, name string) {
	s, err := client.CreateSwimlane(projectID, boardID, name)
	if err != nil {
		log.Fatalf("Error creating swimlane: %v", err)
	}
	fmt.Printf("Created Swimlane: %s (ID: %s)\n", s.Name, s.ID)
}

func UpdateSwimlane(client *sdk.Client, projectID, boardID, swimlaneID, name string) {
	s, err := client.UpdateSwimlane(projectID, boardID, swimlaneID, name)
	if err != nil {
		log.Fatalf("Error updating swimlane: %v", err)
	}
	fmt.Printf("Updated Swimlane: %s (ID: %s)\n", s.Name, s.ID)
}

func DeleteSwimlane(client *sdk.Client, projectID, boardID, swimlaneID string) {
	if err := client.DeleteSwimlane(projectID, boardID, swimlaneID); err != nil {
		log.Fatalf("Error deleting swimlane: %v", err)
	}
	fmt.Printf("Deleted swimlane %s\n", swimlaneID)
}

// ── Kanban Entry CRUD ──────────────────────────────────────────────────────

func CreateKanbanEntry(client *sdk.Client, projectID, boardID, name string, wipLimit int, isSplit bool, dod string) {
	e, err := client.CreateKanbanEntry(projectID, boardID, name, wipLimit, isSplit, dod)
	if err != nil {
		log.Fatalf("Error creating kanban entry: %v", err)
	}
	fmt.Printf("Created Entry: %s (ID: %s)\n", e.Name, e.ID)
}

func UpdateKanbanEntry(client *sdk.Client, projectID, boardID, entryID, name string, wipLimit int, isSplit bool, dod string) {
	e, err := client.UpdateKanbanEntry(projectID, boardID, entryID, name, wipLimit, isSplit, dod)
	if err != nil {
		log.Fatalf("Error updating kanban entry: %v", err)
	}
	fmt.Printf("Updated Entry: %s (ID: %s)\n", e.Name, e.ID)
}

func DeleteKanbanEntry(client *sdk.Client, projectID, boardID, entryID string) {
	if err := client.DeleteKanbanEntry(projectID, boardID, entryID); err != nil {
		log.Fatalf("Error deleting kanban entry: %v", err)
	}
	fmt.Printf("Deleted entry %s\n", entryID)
}

// ── Project Member CRUD ────────────────────────────────────────────────────

func AddProjectMember(client *sdk.Client, projectID, memberID, memberType, roleID string) {
	if projectID == "" || memberID == "" || memberType == "" {
		log.Fatal("Project ID, Member ID, and Member Type are required")
	}
	m, err := client.AddProjectMember(projectID, memberID, memberType, roleID)
	if err != nil {
		log.Fatalf("Error adding project member: %v", err)
	}
	fmt.Printf("Added Member: %s (ID: %s)\n", memberID, m.ID)
}

func RemoveProjectMember(client *sdk.Client, projectID, memberID string) {
	if projectID == "" || memberID == "" {
		log.Fatal("Project ID and Member ID are required")
	}
	err := client.RemoveProjectMember(projectID, memberID)
	if err != nil {
		log.Fatalf("Error removing project member: %v", err)
	}
	fmt.Printf("Removed Member %s from Project %s\n", memberID, projectID)
}

func UpdateProjectMember(client *sdk.Client, projectID, memberID, roleID string) {
	if projectID == "" || memberID == "" {
		log.Fatal("Project ID and Member ID are required")
	}
	m, err := client.UpdateProjectMember(projectID, memberID, roleID)
	if err != nil {
		log.Fatalf("Error updating project member: %v", err)
	}
	fmt.Printf("Updated Member %s (ID: %s)\n", memberID, m.ID)
}

// ── WorkItem Participant ───────────────────────────────────────────────────

func ListWorkItemParticipants(client *sdk.Client, workItemID string) {
	if workItemID == "" {
		log.Fatal("WorkItem ID is required")
	}
	participants, err := client.ListWorkItemParticipants(workItemID)
	if err != nil {
		log.Fatalf("Error listing work item participants: %v", err)
	}
	fmt.Println("Participants:")
	for _, p := range participants {
		fmt.Printf("- ID: %s (Type: %s)\n", p.ID, p.Type)
	}
}

func AddWorkItemParticipant(client *sdk.Client, workItemID, userID string) {
	if workItemID == "" || userID == "" {
		log.Fatal("WorkItem ID and User ID are required")
	}
	p, err := client.AddWorkItemParticipant(workItemID, userID)
	if err != nil {
		log.Fatalf("Error adding work item participant: %v", err)
	}
	fmt.Printf("Added Participant: %s (Type: %s)\n", p.ID, p.Type)
}

func RemoveWorkItemParticipant(client *sdk.Client, workItemID, participantID string) {
	if workItemID == "" || participantID == "" {
		log.Fatal("WorkItem ID and Participant ID are required")
	}
	err := client.RemoveWorkItemParticipant(workItemID, participantID)
	if err != nil {
		log.Fatalf("Error removing work item participant: %v", err)
	}
	fmt.Printf("Removed Participant %s from WorkItem %s\n", participantID, workItemID)
}

// ── WorkItem Relation ──────────────────────────────────────────────────────

func ListWorkItemRelations(client *sdk.Client, workItemID string) {
	if workItemID == "" {
		log.Fatal("WorkItem ID is required")
	}
	relations, err := client.ListWorkItemRelations(workItemID)
	if err != nil {
		log.Fatalf("Error listing work item relations: %v", err)
	}
	fmt.Println("Relations:")
	for _, r := range relations {
		fmt.Printf("- ID: %s (Type: %s)\n", r.ID, r.RelationType)
	}
}

func AddWorkItemRelation(client *sdk.Client, workItemID, targetID, relationType string) {
	if workItemID == "" || targetID == "" || relationType == "" {
		log.Fatal("WorkItem ID, Target ID, and Relation Type are required")
	}
	r, err := client.AddWorkItemRelation(workItemID, targetID, relationType)
	if err != nil {
		log.Fatalf("Error adding work item relation: %v", err)
	}
	fmt.Printf("Added Relation: %s (Type: %s)\n", r.ID, r.RelationType)
}

func RemoveWorkItemRelation(client *sdk.Client, workItemID, relationID string) {
	if workItemID == "" || relationID == "" {
		log.Fatal("WorkItem ID and Relation ID are required")
	}
	err := client.RemoveWorkItemRelation(workItemID, relationID)
	if err != nil {
		log.Fatalf("Error removing work item relation: %v", err)
	}
	fmt.Printf("Removed Relation %s from WorkItem %s\n", relationID, workItemID)
}

// ── WorkItem Transition History ────────────────────────────────────────────

func ListWorkItemTransitionHistories(client *sdk.Client, workItemID string) {
	if workItemID == "" {
		log.Fatal("WorkItem ID is required")
	}
	histories, err := client.ListWorkItemTransitionHistories(workItemID)
	if err != nil {
		log.Fatalf("Error listing work item transition histories: %v", err)
	}
	fmt.Println("Transition Histories:")
	for _, h := range histories {
		from, to := "", ""
		if h.FromState != nil {
			from = h.FromState.Name
		}
		if h.ToState != nil {
			to = h.ToState.Name
		}
		fmt.Printf("- ID: %s  %s → %s  (at: %d)\n", h.ID, from, to, h.CreatedAt)
	}
}

// ── Sprint Sections ────────────────────────────────────────────────────────

func ListSprintSections(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	sections, err := client.ListSprintSections(projectID)
	if err != nil {
		log.Fatalf("Error listing sprint sections: %v", err)
	}
	fmt.Println("Sprint Sections:")
	for _, s := range sections {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func CreateSprintSection(client *sdk.Client, projectID, name string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and name are required")
	}
	s, err := client.CreateSprintSection(projectID, name)
	if err != nil {
		log.Fatalf("Error creating sprint section: %v", err)
	}
	fmt.Printf("Created Section: %s (ID: %s)\n", s.Name, s.ID)
}

func UpdateSprintSection(client *sdk.Client, projectID, sectionID, name string) {
	if projectID == "" || sectionID == "" || name == "" {
		log.Fatal("Project ID, section ID, and name are required")
	}
	s, err := client.UpdateSprintSection(projectID, sectionID, name)
	if err != nil {
		log.Fatalf("Error updating sprint section: %v", err)
	}
	fmt.Printf("Updated Section: %s (ID: %s)\n", s.Name, s.ID)
}

func DeleteSprintSection(client *sdk.Client, projectID, sectionID string) {
	if projectID == "" || sectionID == "" {
		log.Fatal("Project ID and section ID are required")
	}
	err := client.DeleteSprintSection(projectID, sectionID)
	if err != nil {
		log.Fatalf("Error deleting sprint section: %v", err)
	}
	fmt.Printf("Deleted section %s\n", sectionID)
}

func ListSprintCategories(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	cats, err := client.ListSprintCategories(projectID)
	if err != nil {
		log.Fatalf("Error listing sprint categories: %v", err)
	}
	fmt.Println("Sprint Categories:")
	for _, c := range cats {
		fmt.Printf("- %s (ID: %s)\n", c.Name, c.ID)
	}
}

func CreateSprintCategory(client *sdk.Client, projectID, name, sectionID string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and name are required")
	}
	cat, err := client.CreateSprintCategory(projectID, name, sectionID)
	if err != nil {
		log.Fatalf("Error creating sprint category: %v", err)
	}
	fmt.Printf("Created Category: %s (ID: %s)\n", cat.Name, cat.ID)
}

func UpdateSprintCategory(client *sdk.Client, projectID, categoryID, name string) {
	if projectID == "" || categoryID == "" || name == "" {
		log.Fatal("Project ID, category ID, and name are required")
	}
	cat, err := client.UpdateSprintCategory(projectID, categoryID, name)
	if err != nil {
		log.Fatalf("Error updating sprint category: %v", err)
	}
	fmt.Printf("Updated Category: %s (ID: %s)\n", cat.Name, cat.ID)
}

func DeleteSprintCategory(client *sdk.Client, projectID, categoryID string) {
	if projectID == "" || categoryID == "" {
		log.Fatal("Project ID and category ID are required")
	}
	err := client.DeleteSprintCategory(projectID, categoryID)
	if err != nil {
		log.Fatalf("Error deleting sprint category: %v", err)
	}
	fmt.Printf("Deleted category %s\n", categoryID)
}

// ── Version Sections ───────────────────────────────────────────────────────

func ListVersionSections(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	sections, err := client.ListVersionSections(projectID)
	if err != nil {
		log.Fatalf("Error listing version sections: %v", err)
	}
	fmt.Println("Version Sections:")
	for _, s := range sections {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func CreateVersionSection(client *sdk.Client, projectID, name string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and name are required")
	}
	s, err := client.CreateVersionSection(projectID, name)
	if err != nil {
		log.Fatalf("Error creating version section: %v", err)
	}
	fmt.Printf("Created Section: %s (ID: %s)\n", s.Name, s.ID)
}

func UpdateVersionSection(client *sdk.Client, projectID, sectionID, name string) {
	if projectID == "" || sectionID == "" || name == "" {
		log.Fatal("Project ID, section ID, and name are required")
	}
	s, err := client.UpdateVersionSection(projectID, sectionID, name)
	if err != nil {
		log.Fatalf("Error updating version section: %v", err)
	}
	fmt.Printf("Updated Section: %s (ID: %s)\n", s.Name, s.ID)
}

func DeleteVersionSection(client *sdk.Client, projectID, sectionID string) {
	if projectID == "" || sectionID == "" {
		log.Fatal("Project ID and section ID are required")
	}
	err := client.DeleteVersionSection(projectID, sectionID)
	if err != nil {
		log.Fatalf("Error deleting version section: %v", err)
	}
	fmt.Printf("Deleted section %s\n", sectionID)
}

func ListVersionCategories(client *sdk.Client, projectID string) {
	if projectID == "" {
		log.Fatal("Project ID is required")
	}
	cats, err := client.ListVersionCategories(projectID)
	if err != nil {
		log.Fatalf("Error listing version categories: %v", err)
	}
	fmt.Println("Version Categories:")
	for _, c := range cats {
		fmt.Printf("- %s (ID: %s)\n", c.Name, c.ID)
	}
}

func CreateVersionCategory(client *sdk.Client, projectID, name, sectionID string) {
	if projectID == "" || name == "" {
		log.Fatal("Project ID and name are required")
	}
	cat, err := client.CreateVersionCategory(projectID, name, sectionID)
	if err != nil {
		log.Fatalf("Error creating version category: %v", err)
	}
	fmt.Printf("Created Category: %s (ID: %s)\n", cat.Name, cat.ID)
}

func UpdateVersionCategory(client *sdk.Client, projectID, categoryID, name string) {
	if projectID == "" || categoryID == "" || name == "" {
		log.Fatal("Project ID, category ID, and name are required")
	}
	cat, err := client.UpdateVersionCategory(projectID, categoryID, name)
	if err != nil {
		log.Fatalf("Error updating version category: %v", err)
	}
	fmt.Printf("Updated Category: %s (ID: %s)\n", cat.Name, cat.ID)
}

func DeleteVersionCategory(client *sdk.Client, projectID, categoryID string) {
	if projectID == "" || categoryID == "" {
		log.Fatal("Project ID and category ID are required")
	}
	err := client.DeleteVersionCategory(projectID, categoryID)
	if err != nil {
		log.Fatalf("Error deleting version category: %v", err)
	}
	fmt.Printf("Deleted category %s\n", categoryID)
}

func Help(cmd *cobra.Command, args []string) {
	// ctx := cmd.Context()
	cmd.Help()
}
