package cmd

import (
	"log"
	"pingcode-client/internal/app/project"
	"strings"

	"github.com/spf13/cobra"
)

// splitComma splits a comma-separated string into a slice, filtering empty strings.
func splitComma(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "PingCode 项目管理 (Project/Agile) 模块",
	Long:  `用于管理 PingCode 中的敏捷项目信息，支持项目、工作项、迭代、版本、看板的完整 CRUD 操作。`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	projectID   string
	resourceID  string
	name        string
	desc        string
	title       string
	itemType    string
	typeID      string
	status      string
	priority    string
	assignee    string
	startDate   string
	endDate     string
	releaseDate string
	stageID     string
	identifier  string
	sprintID    string
	boardID     string
	entryID     string
	swimlaneID  string
	parentID    string
	tagID       string
	tagIDs      string
	stateIDs    string
	versionIDs  string
	wipLimit      int
	isSplit       bool
	dod           string
	memberID      string
	memberType    string
	roleID        string
	targetID      string
	relType       string
	participantID string
	relationID    string
	userID        string
)

func init() {
	rootCmd.AddCommand(projectCmd)

	// project subcommands
	projectCmd.AddCommand(projectListCmd)
	projectCmd.AddCommand(projectGetCmd)
	projectCmd.AddCommand(projectCreateCmd)
	projectCmd.AddCommand(projectUpdateCmd)
	projectCmd.AddCommand(projectDeleteCmd)

	// workitem subcommands
	projectCmd.AddCommand(workitemCmd)
	workitemCmd.AddCommand(workitemListCmd)
	workitemCmd.AddCommand(workitemGetCmd)
	workitemCmd.AddCommand(workitemCreateCmd)
	workitemCmd.AddCommand(workitemUpdateCmd)
	workitemCmd.AddCommand(workitemDeleteCmd)
	workitemCmd.AddCommand(workitemTagCmd)
	workitemTagCmd.AddCommand(workitemTagAddCmd)
	workitemTagCmd.AddCommand(workitemTagRemoveCmd)

	// iteration subcommands
	projectCmd.AddCommand(iterationCmd)
	iterationCmd.AddCommand(iterationListCmd)
	iterationCmd.AddCommand(iterationGetCmd)
	iterationCmd.AddCommand(iterationCreateCmd)
	iterationCmd.AddCommand(iterationUpdateCmd)
	iterationCmd.AddCommand(iterationDeleteCmd)

	// version subcommands
	projectCmd.AddCommand(versionCmd)
	versionCmd.AddCommand(versionListCmd)
	versionCmd.AddCommand(versionGetCmd)
	versionCmd.AddCommand(versionCreateCmd)
	versionCmd.AddCommand(versionUpdateCmd)
	versionCmd.AddCommand(versionDeleteCmd)

	// kanban subcommands
	projectCmd.AddCommand(kanbanCmd)
	kanbanCmd.AddCommand(kanbanListCmd)
	kanbanCmd.AddCommand(kanbanGetCmd)
	kanbanCmd.AddCommand(kanbanCreateCmd)
	kanbanCmd.AddCommand(kanbanUpdateCmd)
	kanbanCmd.AddCommand(kanbanDeleteCmd)
	kanbanCmd.AddCommand(kanbanEntriesCmd)

	// kanban swimlane subcommands
	kanbanCmd.AddCommand(swimlaneCmd)
	swimlaneCmd.AddCommand(swimlaneListCmd)
	swimlaneCmd.AddCommand(swimlaneCreateCmd)
	swimlaneCmd.AddCommand(swimlaneUpdateCmd)
	swimlaneCmd.AddCommand(swimlaneDeleteCmd)

	// kanban entry subcommands
	kanbanCmd.AddCommand(entryCmd)
	entryCmd.AddCommand(entryCreateCmd)
	entryCmd.AddCommand(entryUpdateCmd)
	entryCmd.AddCommand(entryDeleteCmd)

	// metadata subcommands
	projectCmd.AddCommand(metadataCmd)
	metadataCmd.AddCommand(typesCmd)
	metadataCmd.AddCommand(prioritiesCmd)
	metadataCmd.AddCommand(statusesCmd)
	metadataCmd.AddCommand(severitiesCmd)
	metadataCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsListCmd)
	tagsCmd.AddCommand(tagsCreateCmd)
	metadataCmd.AddCommand(propertiesCmd)
	metadataCmd.AddCommand(membersCmd)

	// member subcommands (project member CRUD)
	projectCmd.AddCommand(memberCmd)
	memberCmd.AddCommand(memberListCmd)
	memberCmd.AddCommand(memberAddCmd)
	memberCmd.AddCommand(memberRemoveCmd)
	memberCmd.AddCommand(memberUpdateCmd)

	// workitem participant subcommands
	workitemCmd.AddCommand(participantCmd)
	participantCmd.AddCommand(participantListCmd)
	participantCmd.AddCommand(participantAddCmd)
	participantCmd.AddCommand(participantRemoveCmd)

	// workitem relation subcommands
	workitemCmd.AddCommand(relationCmd)
	relationCmd.AddCommand(relationListCmd)
	relationCmd.AddCommand(relationAddCmd)
	relationCmd.AddCommand(relationRemoveCmd)

	// workitem transition history
	workitemCmd.AddCommand(workitemHistoriesCmd)

	// Project
	projectGetCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	projectGetCmd.MarkFlagRequired("id")

	projectCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	projectCreateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description of the resource")
	projectCreateCmd.Flags().StringVar(&itemType, "type", "", "Type of the project")
	projectCreateCmd.MarkFlagRequired("name")
	projectCreateCmd.MarkFlagRequired("type")

	projectUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	projectUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	projectUpdateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description of the resource")
	projectUpdateCmd.MarkFlagRequired("id")

	projectDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	projectDeleteCmd.MarkFlagRequired("id")

	// WorkItem
	workitemListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	workitemListCmd.Flags().StringVar(&tagIDs, "tag-ids", "", "Filter by Tag IDs (comma separated)")
	workitemListCmd.Flags().StringVar(&stateIDs, "state-ids", "", "Filter by State IDs (comma separated)")
	workitemListCmd.MarkFlagRequired("project-id")

	workitemGetCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemGetCmd.Flags().StringVarP(&identifier, "identifier", "k", "", "WorkItem Identifier (key)")

	workitemCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	workitemCreateCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the workitem")
	workitemCreateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description of the workitem")
	workitemCreateCmd.Flags().StringVar(&typeID, "type-id", "", "Type ID of the workitem (required)")
	workitemCreateCmd.Flags().StringVar(&status, "state-id", "", "State ID of the workitem")
	workitemCreateCmd.Flags().StringVar(&priority, "priority-id", "", "Priority ID of the workitem")
	workitemCreateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee ID of the workitem")
	workitemCreateCmd.Flags().StringVar(&sprintID, "sprint-id", "", "Sprint ID")
	workitemCreateCmd.Flags().StringVar(&boardID, "board-id", "", "Board ID")
	workitemCreateCmd.Flags().StringVar(&entryID, "entry-id", "", "Entry ID")
	workitemCreateCmd.Flags().StringVar(&swimlaneID, "swimlane-id", "", "Swimlane ID")
	workitemCreateCmd.Flags().StringVar(&parentID, "parent-id", "", "Parent WorkItem ID")
	workitemCreateCmd.Flags().StringVar(&versionIDs, "version-ids", "", "Version IDs (comma separated)")
	workitemCreateCmd.MarkFlagRequired("project-id")
	workitemCreateCmd.MarkFlagRequired("title")
	workitemCreateCmd.MarkFlagRequired("type-id")

	workitemUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemUpdateCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the workitem")
	workitemUpdateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description of the workitem")
	workitemUpdateCmd.Flags().StringVar(&status, "state-id", "", "State ID of the workitem")
	workitemUpdateCmd.Flags().StringVar(&priority, "priority-id", "", "Priority ID of the workitem")
	workitemUpdateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee ID of the workitem")
	workitemUpdateCmd.Flags().StringVar(&sprintID, "sprint-id", "", "Sprint ID")
	workitemUpdateCmd.Flags().StringVar(&boardID, "board-id", "", "Board ID")
	workitemUpdateCmd.Flags().StringVar(&entryID, "entry-id", "", "Entry ID")
	workitemUpdateCmd.Flags().StringVar(&swimlaneID, "swimlane-id", "", "Swimlane ID")
	workitemUpdateCmd.Flags().StringVar(&parentID, "parent-id", "", "Parent WorkItem ID")
	workitemUpdateCmd.Flags().StringVar(&versionIDs, "version-ids", "", "Version IDs (comma separated)")
	workitemUpdateCmd.MarkFlagRequired("id")

	workitemDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemDeleteCmd.MarkFlagRequired("id")

	workitemTagAddCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemTagAddCmd.Flags().StringVarP(&tagID, "tag-id", "t", "", "Tag ID")
	workitemTagAddCmd.MarkFlagRequired("id")
	workitemTagAddCmd.MarkFlagRequired("tag-id")

	workitemTagRemoveCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemTagRemoveCmd.Flags().StringVarP(&tagID, "tag-id", "t", "", "Tag ID")
	workitemTagRemoveCmd.MarkFlagRequired("id")
	workitemTagRemoveCmd.MarkFlagRequired("tag-id")

	// Iteration
	iterationListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	iterationListCmd.MarkFlagRequired("project-id")

	iterationGetCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	iterationGetCmd.MarkFlagRequired("id")

	iterationCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	iterationCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	iterationCreateCmd.Flags().StringVar(&startDate, "start-date", "", "Start date (YYYY-MM-DD)")
	iterationCreateCmd.Flags().StringVar(&endDate, "end-date", "", "End date (YYYY-MM-DD)")
	iterationCreateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee user ID")
	iterationCreateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description")
	iterationCreateCmd.Flags().StringVar(&status, "status", "", "Iteration status")
	iterationCreateCmd.MarkFlagRequired("project-id")
	iterationCreateCmd.MarkFlagRequired("name")

	iterationUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	iterationUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	iterationUpdateCmd.Flags().StringVar(&startDate, "start-date", "", "Start date (YYYY-MM-DD)")
	iterationUpdateCmd.Flags().StringVar(&endDate, "end-date", "", "End date (YYYY-MM-DD)")
	iterationUpdateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee user ID")
	iterationUpdateCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description")
	iterationUpdateCmd.Flags().StringVar(&status, "status", "", "Iteration status")
	iterationUpdateCmd.MarkFlagRequired("id")

	iterationDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	iterationDeleteCmd.MarkFlagRequired("id")

	// Version
	versionListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	versionListCmd.MarkFlagRequired("project-id")

	versionGetCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	versionGetCmd.MarkFlagRequired("id")

	versionCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	versionCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	versionCreateCmd.Flags().StringVar(&startDate, "start-date", "", "Start date (YYYY-MM-DD)")
	versionCreateCmd.Flags().StringVar(&releaseDate, "end-date", "", "Release/end date (YYYY-MM-DD)")
	versionCreateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee user ID")
	versionCreateCmd.Flags().StringVar(&stageID, "stage-id", "", "Stage ID")
	versionCreateCmd.MarkFlagRequired("project-id")
	versionCreateCmd.MarkFlagRequired("name")

	versionUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	versionUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	versionUpdateCmd.Flags().StringVar(&startDate, "start-date", "", "Start date (YYYY-MM-DD)")
	versionUpdateCmd.Flags().StringVar(&releaseDate, "end-date", "", "Release/end date (YYYY-MM-DD)")
	versionUpdateCmd.Flags().StringVar(&assignee, "assignee-id", "", "Assignee user ID")
	versionUpdateCmd.Flags().StringVar(&stageID, "stage-id", "", "Stage ID")
	versionUpdateCmd.MarkFlagRequired("id")

	versionDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	versionDeleteCmd.MarkFlagRequired("id")

	// Kanban
	kanbanListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	kanbanListCmd.MarkFlagRequired("project-id")

	kanbanGetCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	kanbanGetCmd.MarkFlagRequired("id")

	kanbanCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	kanbanCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	kanbanCreateCmd.MarkFlagRequired("project-id")
	kanbanCreateCmd.MarkFlagRequired("name")

	kanbanUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	kanbanUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
	kanbanUpdateCmd.MarkFlagRequired("id")

	kanbanDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	kanbanDeleteCmd.MarkFlagRequired("id")

	kanbanEntriesCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	kanbanEntriesCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	kanbanEntriesCmd.MarkFlagRequired("project-id")
	kanbanEntriesCmd.MarkFlagRequired("board-id")

	// Swimlane flags
	swimlaneListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	swimlaneListCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	swimlaneListCmd.MarkFlagRequired("project-id")
	swimlaneListCmd.MarkFlagRequired("board-id")

	swimlaneCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	swimlaneCreateCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	swimlaneCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Swimlane name")
	swimlaneCreateCmd.MarkFlagRequired("project-id")
	swimlaneCreateCmd.MarkFlagRequired("board-id")
	swimlaneCreateCmd.MarkFlagRequired("name")

	swimlaneUpdateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	swimlaneUpdateCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	swimlaneUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Swimlane ID")
	swimlaneUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "New swimlane name")
	swimlaneUpdateCmd.MarkFlagRequired("project-id")
	swimlaneUpdateCmd.MarkFlagRequired("board-id")
	swimlaneUpdateCmd.MarkFlagRequired("id")

	swimlaneDeleteCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	swimlaneDeleteCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	swimlaneDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Swimlane ID")
	swimlaneDeleteCmd.MarkFlagRequired("project-id")
	swimlaneDeleteCmd.MarkFlagRequired("board-id")
	swimlaneDeleteCmd.MarkFlagRequired("id")

	// Entry flags
	entryCreateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	entryCreateCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	entryCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Entry (column) name")
	entryCreateCmd.Flags().IntVar(&wipLimit, "wip-limit", 0, "WIP limit (0 = unlimited)")
	entryCreateCmd.Flags().BoolVar(&isSplit, "split", false, "Split into in-progress and done")
	entryCreateCmd.Flags().StringVar(&dod, "dod", "", "Definition of done")
	entryCreateCmd.MarkFlagRequired("project-id")
	entryCreateCmd.MarkFlagRequired("board-id")
	entryCreateCmd.MarkFlagRequired("name")

	entryUpdateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	entryUpdateCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	entryUpdateCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Entry ID")
	entryUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "New entry name")
	entryUpdateCmd.Flags().IntVar(&wipLimit, "wip-limit", 0, "WIP limit (0 = unlimited)")
	entryUpdateCmd.Flags().BoolVar(&isSplit, "split", false, "Split into in-progress and done")
	entryUpdateCmd.Flags().StringVar(&dod, "dod", "", "Definition of done")
	entryUpdateCmd.MarkFlagRequired("project-id")
	entryUpdateCmd.MarkFlagRequired("board-id")
	entryUpdateCmd.MarkFlagRequired("id")

	entryDeleteCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	entryDeleteCmd.Flags().StringVarP(&boardID, "board-id", "b", "", "Kanban Board ID")
	entryDeleteCmd.Flags().StringVarP(&resourceID, "id", "i", "", "Entry ID")
	entryDeleteCmd.MarkFlagRequired("project-id")
	entryDeleteCmd.MarkFlagRequired("board-id")
	entryDeleteCmd.MarkFlagRequired("id")

	// Metadata
	prioritiesCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	prioritiesCmd.MarkFlagRequired("project-id")

	// tagsListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	// tagsListCmd.MarkFlagRequired("project-id")

	tagsCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Tag Name")
	tagsCreateCmd.MarkFlagRequired("name")

	propertiesCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	propertiesCmd.Flags().StringVar(&typeID, "type-id", "", "WorkItem Type ID")
	propertiesCmd.MarkFlagRequired("project-id")
	propertiesCmd.MarkFlagRequired("type-id")

	membersCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	membersCmd.MarkFlagRequired("project-id")

	// Member CRUD flags
	memberListCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	memberListCmd.MarkFlagRequired("project-id")

	memberAddCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	memberAddCmd.Flags().StringVar(&memberID, "member-id", "", "Member ID")
	memberAddCmd.Flags().StringVar(&memberType, "member-type", "", "Member type (e.g. user)")
	memberAddCmd.Flags().StringVar(&roleID, "role-id", "", "Role ID (optional)")
	memberAddCmd.MarkFlagRequired("project-id")
	memberAddCmd.MarkFlagRequired("member-id")
	memberAddCmd.MarkFlagRequired("member-type")

	memberRemoveCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	memberRemoveCmd.Flags().StringVar(&memberID, "member-id", "", "Member ID")
	memberRemoveCmd.MarkFlagRequired("project-id")
	memberRemoveCmd.MarkFlagRequired("member-id")

	memberUpdateCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "PingCode Project ID")
	memberUpdateCmd.Flags().StringVar(&memberID, "member-id", "", "Member ID")
	memberUpdateCmd.Flags().StringVar(&roleID, "role-id", "", "New Role ID")
	memberUpdateCmd.MarkFlagRequired("project-id")
	memberUpdateCmd.MarkFlagRequired("member-id")
	memberUpdateCmd.MarkFlagRequired("role-id")

	// Participant flags
	participantListCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	participantListCmd.MarkFlagRequired("id")

	participantAddCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	participantAddCmd.Flags().StringVar(&userID, "user-id", "", "User ID to add as participant")
	participantAddCmd.MarkFlagRequired("id")
	participantAddCmd.MarkFlagRequired("user-id")

	participantRemoveCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	participantRemoveCmd.Flags().StringVar(&participantID, "participant-id", "", "Participant ID to remove")
	participantRemoveCmd.MarkFlagRequired("id")
	participantRemoveCmd.MarkFlagRequired("participant-id")

	// Relation flags
	relationListCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	relationListCmd.MarkFlagRequired("id")

	relationAddCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	relationAddCmd.Flags().StringVar(&targetID, "target-id", "", "Target WorkItem ID")
	relationAddCmd.Flags().StringVar(&relType, "type", "", "Relation type")
	relationAddCmd.MarkFlagRequired("id")
	relationAddCmd.MarkFlagRequired("target-id")
	relationAddCmd.MarkFlagRequired("type")

	relationRemoveCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	relationRemoveCmd.Flags().StringVar(&relationID, "relation-id", "", "Relation ID to remove")
	relationRemoveCmd.MarkFlagRequired("id")
	relationRemoveCmd.MarkFlagRequired("relation-id")

	// Transition history flags
	workitemHistoriesCmd.Flags().StringVarP(&resourceID, "id", "i", "", "WorkItem ID")
	workitemHistoriesCmd.MarkFlagRequired("id")
}

// Resource Commands
var workitemCmd = &cobra.Command{Use: "workitem", Short: "工作项管理"}
var iterationCmd = &cobra.Command{Use: "iteration", Short: "迭代 (Scrum) 管理"}
var versionCmd = &cobra.Command{Use: "version", Short: "发布版本管理"}
var kanbanCmd = &cobra.Command{Use: "kanban", Short: "看板管理"}
var metadataCmd = &cobra.Command{Use: "metadata", Short: "元数据查询"}

// Project CRUD
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出项目",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListProjects(c)
	},
}
var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取项目详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.GetProject(c, resourceID)
	},
}
var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建项目",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateProject(c, name, desc, itemType)
	},
}
var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新项目",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.UpdateProject(c, resourceID, name, desc)
	},
}
var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除项目",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.DeleteProject(c, resourceID)
	},
}

// WorkItem CRUD
var workitemListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出工作项",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListWorkItems(c, projectID, tagIDs, stateIDs)
	},
}
var workitemGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取工作项详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.GetWorkItem(c, resourceID, identifier)
	},
}
var workitemCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建工作项",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateWorkItem(c, projectID, title, desc, itemType, typeID, status, priority, assignee, sprintID, boardID, entryID, swimlaneID, parentID, splitComma(versionIDs))
	},
}
var workitemUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新工作项",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.UpdateWorkItem(c, resourceID, title, desc, status, priority, assignee, sprintID, boardID, entryID, swimlaneID, parentID, splitComma(versionIDs))
	},
}
var workitemDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除工作项",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.DeleteWorkItem(c, resourceID)
	},
}

var workitemTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "工作项标签管理",
}

var workitemTagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "向工作项中添加一个标签",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.AddWorkItemTag(c, resourceID, tagID)
	},
}

var workitemTagRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "在工作项中移除一个标签",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.RemoveWorkItemTag(c, resourceID, tagID)
	},
}

// Iteration CRUD
var iterationListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出迭代",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListIterations(c, projectID)
	},
}
var iterationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取迭代详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.GetIteration(c, resourceID)
	},
}
var iterationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建迭代",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateIteration(c, projectID, name, startDate, endDate, assignee, desc, status)
	},
}
var iterationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新迭代",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.UpdateIteration(c, resourceID, name, startDate, endDate, assignee, desc, status)
	},
}
var iterationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除迭代",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.DeleteIteration(c, resourceID)
	},
}

// Version CRUD
var versionListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出版本",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListVersions(c, projectID)
	},
}
var versionGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取版本详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.GetVersion(c, resourceID)
	},
}
var versionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建版本",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateVersion(c, projectID, name, startDate, releaseDate, assignee, stageID)
	},
}
var versionUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新版本",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.UpdateVersion(c, resourceID, name, startDate, releaseDate, assignee, stageID)
	},
}
var versionDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除版本",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.DeleteVersion(c, resourceID)
	},
}

// Kanban CRUD
var kanbanListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出看板",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListKanbans(c, projectID)
	},
}
var kanbanGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取看板详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.GetKanban(c, resourceID)
	},
}
var kanbanCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建看板",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateKanban(c, projectID, name)
	},
}
var kanbanUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新看板",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.UpdateKanban(c, resourceID, name)
	},
}
var kanbanDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除看板",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.DeleteKanban(c, resourceID)
	},
}
var kanbanEntriesCmd = &cobra.Command{
	Use:   "entries",
	Short: "列出看板栏 (Columns)",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListKanbanEntries(c, projectID, boardID)
	},
}
var severitiesCmd = &cobra.Command{
	Use:   "severities",
	Short: "列出项目管理中的严重程度定义",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListSeverities(c)
	},
}

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "列出项目管理中的工作项类型",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListWorkItemTypes(c)
	},
}

var prioritiesCmd = &cobra.Command{
	Use:   "priorities",
	Short: "列出项目管理中的优先级定义",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListPriorities(c, projectID)
	},
}

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "工作项标签管理",
}

var tagsListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出工作项标签列表",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListTags(c)
	},
}

var tagsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个新的工作项标签",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.CreateWorkItemTag(c, name)
	},
}

var propertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "列出工作项属性列表",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListProperties(c, projectID, typeID)
	},
}

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "列出项目成员列表",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListProjectMembers(c, projectID)
	},
}

var statusesCmd = &cobra.Command{
	Use:   "statuses",
	Short: "列出项目管理中的状态定义",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		project.ListStatuses(c)
	},
}

// ── Swimlane commands ──────────────────────────────────────────────────────

var swimlaneCmd = &cobra.Command{Use: "swimlane", Short: "泳道管理"}

var swimlaneListCmd = &cobra.Command{
Use:   "list",
Short: "列出泳道",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.ListSwimlanes(c, projectID, boardID)
},
}

var swimlaneCreateCmd = &cobra.Command{
Use:   "create",
Short: "创建泳道",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.CreateSwimlane(c, projectID, boardID, name)
},
}

var swimlaneUpdateCmd = &cobra.Command{
Use:   "update",
Short: "更新泳道",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.UpdateSwimlane(c, projectID, boardID, resourceID, name)
},
}

var swimlaneDeleteCmd = &cobra.Command{
Use:   "delete",
Short: "删除泳道",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.DeleteSwimlane(c, projectID, boardID, resourceID)
},
}

// ── Entry commands ─────────────────────────────────────────────────────────

var entryCmd = &cobra.Command{Use: "entry", Short: "看板栏 (Column) 管理"}

var entryCreateCmd = &cobra.Command{
Use:   "create",
Short: "创建看板栏",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.CreateKanbanEntry(c, projectID, boardID, name, wipLimit, isSplit, dod)
},
}

var entryUpdateCmd = &cobra.Command{
Use:   "update",
Short: "更新看板栏",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.UpdateKanbanEntry(c, projectID, boardID, resourceID, name, wipLimit, isSplit, dod)
},
}

var entryDeleteCmd = &cobra.Command{
Use:   "delete",
Short: "删除看板栏",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.DeleteKanbanEntry(c, projectID, boardID, resourceID)
},
}

// ── Member commands ────────────────────────────────────────────────────────

var memberCmd = &cobra.Command{Use: "member", Short: "项目成员管理"}

var memberListCmd = &cobra.Command{
Use:   "list",
Short: "列出项目成员",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.ListProjectMembers(c, projectID)
},
}

var memberAddCmd = &cobra.Command{
Use:   "add",
Short: "添加项目成员",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.AddProjectMember(c, projectID, memberID, memberType, roleID)
},
}

var memberRemoveCmd = &cobra.Command{
Use:   "remove",
Short: "移除项目成员",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.RemoveProjectMember(c, projectID, memberID)
},
}

var memberUpdateCmd = &cobra.Command{
Use:   "update",
Short: "更新项目成员角色",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.UpdateProjectMember(c, projectID, memberID, roleID)
},
}

// ── WorkItem Participant commands ──────────────────────────────────────────

var participantCmd = &cobra.Command{Use: "participant", Short: "工作项参与者管理"}

var participantListCmd = &cobra.Command{
Use:   "list",
Short: "列出工作项参与者",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.ListWorkItemParticipants(c, resourceID)
},
}

var participantAddCmd = &cobra.Command{
Use:   "add",
Short: "添加工作项参与者",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.AddWorkItemParticipant(c, resourceID, userID)
},
}

var participantRemoveCmd = &cobra.Command{
Use:   "remove",
Short: "移除工作项参与者",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.RemoveWorkItemParticipant(c, resourceID, participantID)
},
}

// ── WorkItem Relation commands ─────────────────────────────────────────────

var relationCmd = &cobra.Command{Use: "relation", Short: "工作项关联管理"}

var relationListCmd = &cobra.Command{
Use:   "list",
Short: "列出工作项关联",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.ListWorkItemRelations(c, resourceID)
},
}

var relationAddCmd = &cobra.Command{
Use:   "add",
Short: "添加工作项关联",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.AddWorkItemRelation(c, resourceID, targetID, relType)
},
}

var relationRemoveCmd = &cobra.Command{
Use:   "remove",
Short: "移除工作项关联",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.RemoveWorkItemRelation(c, resourceID, relationID)
},
}

// ── WorkItem Transition History ────────────────────────────────────────────

var workitemHistoriesCmd = &cobra.Command{
Use:   "histories",
Short: "列出工作项状态变更历史",
Run: func(cmd *cobra.Command, args []string) {
c, err := GetClient()
if err != nil {
log.Fatal(err)
}
project.ListWorkItemTransitionHistories(c, resourceID)
},
}
