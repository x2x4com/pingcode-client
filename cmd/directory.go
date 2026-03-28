package cmd

import (
	"log"
	"pingcode-client/internal/app/directory"
	"pingcode-client/internal/pkg/output"

	"github.com/spf13/cobra"
)

var (
	dirUserID      string
	dirName        string
	dirDisplayName string
	dirEmail       string
	dirMobile      string
	dirStatus      string
	dirPassword    string
	dirDeptID      string
	dirJobID       string
	dirEmpNum      string
	dirHeadID      string
	dirParentID    string
	dirKeywords    string
	dirDeptIDs     string
)

var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "PingCode 通讯录模块",
	Long:  `用于管理 PingCode 通讯录中的用户和部门。`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// ── Directory User commands ────────────────────────────────────────────────

var dirUserCmd = &cobra.Command{Use: "user", Short: "用户管理"}

var dirUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出通讯录用户",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.ListUsers(c, dirKeywords, dirName, dirDeptIDs)
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

var dirUserCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建通讯录用户",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.CreateUser(c, dirName, dirDisplayName, dirEmail, dirMobile, dirPassword, dirDeptID, dirJobID, dirEmpNum)
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

var dirUserUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新通讯录用户",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.UpdateUser(c, dirUserID, dirName, dirDisplayName, dirEmail, dirMobile, dirStatus, dirDeptID, dirJobID, dirEmpNum)
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

// ── Department commands ────────────────────────────────────────────────────

var dirDeptCmd = &cobra.Command{Use: "department", Short: "部门管理"}

var dirDeptListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出部门",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.ListDepartments(c)
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

var dirDeptCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建部门",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.CreateDepartment(c, dirName, dirParentID, dirHeadID)
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

var dirDeptUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新部门",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := directory.UpdateDepartment(c, dirUserID, dirName, dirParentID, dirHeadID)
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

var dirDeptDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除部门",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		err = directory.DeleteDepartment(c, dirUserID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		output.FormatAndPrint(map[string]string{"message": "Deleted department", "id": dirUserID}, opts)
	},
}

func init() {
	rootCmd.AddCommand(directoryCmd)

	// user subcommands
	directoryCmd.AddCommand(dirUserCmd)
	dirUserCmd.AddCommand(dirUserListCmd)
	dirUserCmd.AddCommand(dirUserCreateCmd)
	dirUserCmd.AddCommand(dirUserUpdateCmd)

	// department subcommands
	directoryCmd.AddCommand(dirDeptCmd)
	dirDeptCmd.AddCommand(dirDeptListCmd)
	dirDeptCmd.AddCommand(dirDeptCreateCmd)
	dirDeptCmd.AddCommand(dirDeptUpdateCmd)
	dirDeptCmd.AddCommand(dirDeptDeleteCmd)

	// User list flags
	dirUserListCmd.Flags().StringVar(&dirKeywords, "keywords", "", "Keywords to search users")
	dirUserListCmd.Flags().StringVar(&dirName, "name", "", "Filter by user name")
	dirUserListCmd.Flags().StringVar(&dirDeptIDs, "department-ids", "", "Filter by department IDs (comma separated)")

	// User create flags
	dirUserCreateCmd.Flags().StringVar(&dirName, "name", "", "User name (login name)")
	dirUserCreateCmd.Flags().StringVar(&dirDisplayName, "display-name", "", "User display name")
	dirUserCreateCmd.Flags().StringVar(&dirEmail, "email", "", "User email")
	dirUserCreateCmd.Flags().StringVar(&dirMobile, "mobile", "", "User mobile")
	dirUserCreateCmd.Flags().StringVar(&dirPassword, "password", "", "User password")
	dirUserCreateCmd.Flags().StringVar(&dirDeptID, "department-id", "", "Department ID")
	dirUserCreateCmd.Flags().StringVar(&dirJobID, "job-id", "", "Job ID")
	dirUserCreateCmd.Flags().StringVar(&dirEmpNum, "employee-number", "", "Employee number")
	dirUserCreateCmd.MarkFlagRequired("name")
	dirUserCreateCmd.MarkFlagRequired("display-name")

	// User update flags
	dirUserUpdateCmd.Flags().StringVar(&dirUserID, "id", "", "User ID")
	dirUserUpdateCmd.Flags().StringVar(&dirName, "name", "", "User name")
	dirUserUpdateCmd.Flags().StringVar(&dirDisplayName, "display-name", "", "User display name")
	dirUserUpdateCmd.Flags().StringVar(&dirEmail, "email", "", "User email")
	dirUserUpdateCmd.Flags().StringVar(&dirMobile, "mobile", "", "User mobile")
	dirUserUpdateCmd.Flags().StringVar(&dirStatus, "status", "", "User status")
	dirUserUpdateCmd.Flags().StringVar(&dirDeptID, "department-id", "", "Department ID")
	dirUserUpdateCmd.Flags().StringVar(&dirJobID, "job-id", "", "Job ID")
	dirUserUpdateCmd.Flags().StringVar(&dirEmpNum, "employee-number", "", "Employee number")
	dirUserUpdateCmd.MarkFlagRequired("id")

	// Department create flags
	dirDeptCreateCmd.Flags().StringVar(&dirName, "name", "", "Department name")
	dirDeptCreateCmd.Flags().StringVar(&dirParentID, "parent-id", "", "Parent department ID")
	dirDeptCreateCmd.Flags().StringVar(&dirHeadID, "head-id", "", "Head user ID")
	dirDeptCreateCmd.MarkFlagRequired("name")

	// Department update flags
	dirDeptUpdateCmd.Flags().StringVar(&dirUserID, "id", "", "Department ID")
	dirDeptUpdateCmd.Flags().StringVar(&dirName, "name", "", "Department name")
	dirDeptUpdateCmd.Flags().StringVar(&dirParentID, "parent-id", "", "Parent department ID")
	dirDeptUpdateCmd.Flags().StringVar(&dirHeadID, "head-id", "", "Head user ID")
	dirDeptUpdateCmd.MarkFlagRequired("id")

	// Department delete flags
	dirDeptDeleteCmd.Flags().StringVar(&dirUserID, "id", "", "Department ID")
	dirDeptDeleteCmd.MarkFlagRequired("id")
}
