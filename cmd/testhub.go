package cmd

import (
	"log"
	"pingcode-client/internal/app/testhub"
	"pingcode-client/internal/pkg/output"

	"github.com/spf13/cobra"
)

var (
	thLibraryID     string
	thMemberID      string
	thMemberType    string
	thRoleID        string
	thSuiteID       string
	thCaseID        string
	thPlanID        string
	thName          string
	thIdentifier    string
	thScopeType     string
	thScopeID       string
	thVisibility    string
	thDesc          string
	thTitle         string
	thTypeID        string
	thAssigneeID    string
	thStartAt       int64
	thEndAt         int64
	thExecutorID    string
	thMaintenanceID string
	thPrecondition  string
)

var testhubCmd = &cobra.Command{
	Use:   "testhub",
	Short: "PingCode 测试管理 (TestHub) 模块",
	Long:  `用于管理 PingCode TestHub 测试库、用例、计划和运行，支持完整的 CRUD 操作。`,
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

// library commands

var thLibraryCmd = &cobra.Command{
	Use:   "library",
	Short: "管理测试库",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thLibraryListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有测试库",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.ListLibraries(c)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMapSlice(data), opts)
		}
	},
}

var thLibraryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建测试库",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.CreateLibrary(c, thName, thIdentifier, thScopeType, thScopeID, thVisibility, thDesc)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thLibraryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新测试库",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.UpdateLibrary(c, thLibraryID, thName, thVisibility, thDesc)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

// library member commands

var thMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "管理测试库成员",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thMemberListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出测试库成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.ListLibraryMembers(c, thLibraryID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMapSlice(data), opts)
		}
	},
}

var thMemberAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加测试库成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.AddLibraryMember(c, thLibraryID, thMemberID, thMemberType, thRoleID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thMemberRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "移除测试库成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		err = testhub.RemoveLibraryMember(c, thLibraryID, thMemberID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		output.FormatAndPrint(map[string]string{"message": "Removed member", "member_id": thMemberID}, opts)
	},
}

// suite commands

var thSuiteCmd = &cobra.Command{
	Use:   "suite",
	Short: "管理测试套件",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thSuiteListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出测试套件",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.ListSuites(c, thLibraryID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMapSlice(data), opts)
		}
	},
}

var thSuiteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建测试套件",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.CreateSuite(c, thLibraryID, thName)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thSuiteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除测试套件",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		err = testhub.DeleteSuite(c, thLibraryID, thSuiteID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		output.FormatAndPrint(map[string]string{"message": "Deleted suite", "suite_id": thSuiteID}, opts)
	},
}

// case commands

var thCaseCmd = &cobra.Command{
	Use:   "case",
	Short: "管理测试用例",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thCaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出测试用例",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.ListCases(c, thLibraryID, thSuiteID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMapSlice(data), opts)
		}
	},
}

var thCaseGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取测试用例详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.GetCase(c, thCaseID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thCaseCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建测试用例",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.CreateCase(c, thLibraryID, thTitle, thSuiteID, thTypeID, thMaintenanceID, thDesc, thPrecondition)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thCaseUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新测试用例",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.UpdateCase(c, thCaseID, thTitle, thSuiteID, thTypeID, thMaintenanceID, thDesc, thPrecondition)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

var thCaseDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除测试用例",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		err = testhub.DeleteCase(c, thCaseID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		output.FormatAndPrint(map[string]string{"message": "Deleted case", "case_id": thCaseID}, opts)
	},
}

// plan commands

var thPlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "管理测试计划",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thPlanListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出测试计划",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.ListPlans(c, thLibraryID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMapSlice(data), opts)
		}
	},
}

var thPlanCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建测试计划",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.CreatePlan(c, thLibraryID, thName, thTypeID, thAssigneeID, thStartAt, thEndAt)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

// run commands

var thRunCmd = &cobra.Command{
	Use:   "run",
	Short: "管理测试运行",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var thRunCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建测试运行",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		data, err := testhub.CreateRun(c, thLibraryID, thPlanID, thCaseID, thExecutorID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			output.FormatAndPrint(data, opts)
		} else {
			output.FormatAndPrint(toMap(data), opts)
		}
	},
}

func init() {
	rootCmd.AddCommand(testhubCmd)

	// library subcommands
	testhubCmd.AddCommand(thLibraryCmd)
	thLibraryCmd.AddCommand(thLibraryListCmd)
	thLibraryCmd.AddCommand(thLibraryCreateCmd)
	thLibraryCmd.AddCommand(thLibraryUpdateCmd)
	thLibraryCmd.AddCommand(thMemberCmd)
	thMemberCmd.AddCommand(thMemberListCmd)
	thMemberCmd.AddCommand(thMemberAddCmd)
	thMemberCmd.AddCommand(thMemberRemoveCmd)
	thLibraryCmd.AddCommand(thSuiteCmd)
	thSuiteCmd.AddCommand(thSuiteListCmd)
	thSuiteCmd.AddCommand(thSuiteCreateCmd)
	thSuiteCmd.AddCommand(thSuiteDeleteCmd)

	// case subcommands
	testhubCmd.AddCommand(thCaseCmd)
	thCaseCmd.AddCommand(thCaseListCmd)
	thCaseCmd.AddCommand(thCaseGetCmd)
	thCaseCmd.AddCommand(thCaseCreateCmd)
	thCaseCmd.AddCommand(thCaseUpdateCmd)
	thCaseCmd.AddCommand(thCaseDeleteCmd)

	// plan subcommands
	testhubCmd.AddCommand(thPlanCmd)
	thPlanCmd.AddCommand(thPlanListCmd)
	thPlanCmd.AddCommand(thPlanCreateCmd)

	// run subcommands
	testhubCmd.AddCommand(thRunCmd)
	thRunCmd.AddCommand(thRunCreateCmd)

	// library flags
	thLibraryCreateCmd.Flags().StringVar(&thName, "name", "", "Library name")
	thLibraryCreateCmd.Flags().StringVar(&thIdentifier, "identifier", "", "Library identifier")
	thLibraryCreateCmd.Flags().StringVar(&thScopeType, "scope-type", "", "Scope type")
	thLibraryCreateCmd.Flags().StringVar(&thScopeID, "scope-id", "", "Scope ID")
	thLibraryCreateCmd.Flags().StringVar(&thVisibility, "visibility", "", "Visibility")
	thLibraryCreateCmd.Flags().StringVar(&thDesc, "desc", "", "Description")
	thLibraryCreateCmd.MarkFlagRequired("name")
	thLibraryCreateCmd.MarkFlagRequired("identifier")

	thLibraryUpdateCmd.Flags().StringVar(&thLibraryID, "id", "", "Library ID")
	thLibraryUpdateCmd.Flags().StringVar(&thName, "name", "", "Library name")
	thLibraryUpdateCmd.Flags().StringVar(&thVisibility, "visibility", "", "Visibility")
	thLibraryUpdateCmd.Flags().StringVar(&thDesc, "desc", "", "Description")
	thLibraryUpdateCmd.MarkFlagRequired("id")

	// member flags
	thMemberListCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thMemberListCmd.MarkFlagRequired("library-id")

	thMemberAddCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thMemberAddCmd.Flags().StringVar(&thMemberID, "member-id", "", "Member ID")
	thMemberAddCmd.Flags().StringVar(&thMemberType, "member-type", "", "Member type")
	thMemberAddCmd.Flags().StringVar(&thRoleID, "role-id", "", "Role ID")
	thMemberAddCmd.MarkFlagRequired("library-id")
	thMemberAddCmd.MarkFlagRequired("member-id")
	thMemberAddCmd.MarkFlagRequired("member-type")

	thMemberRemoveCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thMemberRemoveCmd.Flags().StringVar(&thMemberID, "member-id", "", "Member ID")
	thMemberRemoveCmd.MarkFlagRequired("library-id")
	thMemberRemoveCmd.MarkFlagRequired("member-id")

	// suite flags
	thSuiteListCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thSuiteListCmd.MarkFlagRequired("library-id")

	thSuiteCreateCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thSuiteCreateCmd.Flags().StringVar(&thName, "name", "", "Suite name")
	thSuiteCreateCmd.MarkFlagRequired("library-id")
	thSuiteCreateCmd.MarkFlagRequired("name")

	thSuiteDeleteCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thSuiteDeleteCmd.Flags().StringVar(&thSuiteID, "suite-id", "", "Suite ID")
	thSuiteDeleteCmd.MarkFlagRequired("library-id")
	thSuiteDeleteCmd.MarkFlagRequired("suite-id")

	// case flags
	thCaseListCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thCaseListCmd.Flags().StringVar(&thSuiteID, "suite-id", "", "Filter by Suite ID")
	thCaseListCmd.MarkFlagRequired("library-id")

	thCaseGetCmd.Flags().StringVar(&thCaseID, "id", "", "Case ID")
	thCaseGetCmd.MarkFlagRequired("id")

	thCaseCreateCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thCaseCreateCmd.Flags().StringVar(&thTitle, "title", "", "Case title")
	thCaseCreateCmd.Flags().StringVar(&thSuiteID, "suite-id", "", "Suite ID")
	thCaseCreateCmd.Flags().StringVar(&thTypeID, "type-id", "", "Type ID")
	thCaseCreateCmd.Flags().StringVar(&thMaintenanceID, "maintenance-id", "", "Maintenance ID")
	thCaseCreateCmd.Flags().StringVar(&thDesc, "desc", "", "Description")
	thCaseCreateCmd.Flags().StringVar(&thPrecondition, "precondition", "", "Precondition")
	thCaseCreateCmd.MarkFlagRequired("library-id")
	thCaseCreateCmd.MarkFlagRequired("title")

	thCaseUpdateCmd.Flags().StringVar(&thCaseID, "id", "", "Case ID")
	thCaseUpdateCmd.Flags().StringVar(&thTitle, "title", "", "Case title")
	thCaseUpdateCmd.Flags().StringVar(&thSuiteID, "suite-id", "", "Suite ID")
	thCaseUpdateCmd.Flags().StringVar(&thTypeID, "type-id", "", "Type ID")
	thCaseUpdateCmd.Flags().StringVar(&thMaintenanceID, "maintenance-id", "", "Maintenance ID")
	thCaseUpdateCmd.Flags().StringVar(&thDesc, "desc", "", "Description")
	thCaseUpdateCmd.Flags().StringVar(&thPrecondition, "precondition", "", "Precondition")
	thCaseUpdateCmd.MarkFlagRequired("id")

	thCaseDeleteCmd.Flags().StringVar(&thCaseID, "id", "", "Case ID")
	thCaseDeleteCmd.MarkFlagRequired("id")

	// plan flags
	thPlanListCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thPlanListCmd.MarkFlagRequired("library-id")

	thPlanCreateCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thPlanCreateCmd.Flags().StringVar(&thName, "name", "", "Plan name")
	thPlanCreateCmd.Flags().StringVar(&thTypeID, "type-id", "", "Type ID")
	thPlanCreateCmd.Flags().StringVar(&thAssigneeID, "assignee-id", "", "Assignee ID")
	thPlanCreateCmd.Flags().Int64Var(&thStartAt, "start-at", 0, "Start timestamp (unix ms)")
	thPlanCreateCmd.Flags().Int64Var(&thEndAt, "end-at", 0, "End timestamp (unix ms)")
	thPlanCreateCmd.MarkFlagRequired("library-id")
	thPlanCreateCmd.MarkFlagRequired("name")
	thPlanCreateCmd.MarkFlagRequired("type-id")
	thPlanCreateCmd.MarkFlagRequired("assignee-id")
	thPlanCreateCmd.MarkFlagRequired("start-at")
	thPlanCreateCmd.MarkFlagRequired("end-at")

	// run flags
	thRunCreateCmd.Flags().StringVar(&thLibraryID, "library-id", "", "Library ID")
	thRunCreateCmd.Flags().StringVar(&thPlanID, "plan-id", "", "Plan ID")
	thRunCreateCmd.Flags().StringVar(&thCaseID, "case-id", "", "Case ID")
	thRunCreateCmd.Flags().StringVar(&thExecutorID, "executor-id", "", "Executor ID")
	thRunCreateCmd.MarkFlagRequired("library-id")
	thRunCreateCmd.MarkFlagRequired("plan-id")
	thRunCreateCmd.MarkFlagRequired("case-id")
}
