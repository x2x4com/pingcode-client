package cmd

import (
	"log"
	"pingcode-client/internal/app/wiki"

	"github.com/spf13/cobra"
)

var (
	wikiSpaceID     string
	wikiMemberID    string
	wikiMemberType  string
	wikiRoleID      string
	wikiPageID      string
	wikiSpaceIDFlag string
	wikiParentID    string
	wikiPageType    string
	wikiName        string
	wikiIdentifier  string
	wikiScopeType   string
	wikiScopeID     string
	wikiVisibility  string
	wikiDesc        string
	wikiTitle       string
	wikiContent     string
	wikiFormat      string
)

var wikiCmd = &cobra.Command{
	Use:   "wiki",
	Short: "PingCode 知识管理 (Wiki) 模块",
	Long:  `用于管理 PingCode Wiki 空间和页面，支持完整的 CRUD 操作。`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// space commands
var wikiSpaceCmd = &cobra.Command{
	Use:   "space",
	Short: "管理 Wiki 空间",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var wikiSpaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有 Wiki 空间",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.ListSpaces(c)
	},
}

var wikiSpaceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取 Wiki 空间详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.GetSpace(c, wikiSpaceID)
	},
}

var wikiSpaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建 Wiki 空间",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.CreateSpace(c, wikiName, wikiIdentifier, wikiScopeType, wikiScopeID, wikiVisibility, wikiDesc)
	},
}

var wikiSpaceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新 Wiki 空间",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.UpdateSpace(c, wikiSpaceID, wikiName, wikiVisibility, wikiDesc)
	},
}

var wikiSpaceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除 Wiki 空间",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.DeleteSpace(c, wikiSpaceID)
	},
}

// space member commands
var wikiSpaceMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "管理 Wiki 空间成员",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var wikiSpaceMemberListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出 Wiki 空间成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.ListSpaceMembers(c, wikiSpaceID)
	},
}

var wikiSpaceMemberAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加 Wiki 空间成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.AddSpaceMember(c, wikiSpaceID, wikiMemberID, wikiMemberType, wikiRoleID)
	},
}

var wikiSpaceMemberRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "移除 Wiki 空间成员",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.RemoveSpaceMember(c, wikiSpaceID, wikiMemberID)
	},
}

// page commands
var wikiPageCmd = &cobra.Command{
	Use:   "page",
	Short: "管理 Wiki 页面",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var wikiPageListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出 Wiki 页面",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.ListPages(c, wikiSpaceIDFlag)
	},
}

var wikiPageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取 Wiki 页面详情",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.GetPage(c, wikiPageID)
	},
}

var wikiPageCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建 Wiki 页面",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.CreatePage(c, wikiSpaceIDFlag, wikiParentID, wikiTitle, wikiPageType)
	},
}

var wikiPageUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新 Wiki 页面",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.UpdatePage(c, wikiPageID, wikiTitle)
	},
}

var wikiPageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除 Wiki 页面",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.DeletePage(c, wikiPageID)
	},
}

// page content commands
var wikiPageContentCmd = &cobra.Command{
	Use:   "content",
	Short: "管理 Wiki 页面内容",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var wikiPageContentGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取 Wiki 页面内容",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.GetPageContent(c, wikiPageID)
	},
}

var wikiPageContentUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新 Wiki 页面内容",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.UpdatePageContent(c, wikiPageID, wikiContent, wikiFormat)
	},
}

var wikiPageVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "列出 Wiki 页面版本",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := GetClient()
		if err != nil {
			log.Fatalf("Failed to get client: %v", err)
		}
		wiki.ListPageVersions(c, wikiPageID)
	},
}

func init() {
	rootCmd.AddCommand(wikiCmd)

	// space subcommands
	wikiCmd.AddCommand(wikiSpaceCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceListCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceGetCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceCreateCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceUpdateCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceDeleteCmd)
	wikiSpaceCmd.AddCommand(wikiSpaceMemberCmd)
	wikiSpaceMemberCmd.AddCommand(wikiSpaceMemberListCmd)
	wikiSpaceMemberCmd.AddCommand(wikiSpaceMemberAddCmd)
	wikiSpaceMemberCmd.AddCommand(wikiSpaceMemberRemoveCmd)

	// page subcommands
	wikiCmd.AddCommand(wikiPageCmd)
	wikiPageCmd.AddCommand(wikiPageListCmd)
	wikiPageCmd.AddCommand(wikiPageGetCmd)
	wikiPageCmd.AddCommand(wikiPageCreateCmd)
	wikiPageCmd.AddCommand(wikiPageUpdateCmd)
	wikiPageCmd.AddCommand(wikiPageDeleteCmd)
	wikiPageCmd.AddCommand(wikiPageContentCmd)
	wikiPageContentCmd.AddCommand(wikiPageContentGetCmd)
	wikiPageContentCmd.AddCommand(wikiPageContentUpdateCmd)
	wikiPageCmd.AddCommand(wikiPageVersionsCmd)

	// space flags
	wikiSpaceGetCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceGetCmd.MarkFlagRequired("id")

	wikiSpaceCreateCmd.Flags().StringVarP(&wikiName, "name", "n", "", "Space name")
	wikiSpaceCreateCmd.Flags().StringVar(&wikiIdentifier, "identifier", "", "Space identifier")
	wikiSpaceCreateCmd.Flags().StringVar(&wikiScopeType, "scope-type", "", "Scope type")
	wikiSpaceCreateCmd.Flags().StringVar(&wikiScopeID, "scope-id", "", "Scope ID")
	wikiSpaceCreateCmd.Flags().StringVar(&wikiVisibility, "visibility", "", "Visibility")
	wikiSpaceCreateCmd.Flags().StringVar(&wikiDesc, "desc", "", "Description")
	wikiSpaceCreateCmd.MarkFlagRequired("name")
	wikiSpaceCreateCmd.MarkFlagRequired("identifier")
	wikiSpaceCreateCmd.MarkFlagRequired("scope-type")

	wikiSpaceUpdateCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceUpdateCmd.Flags().StringVarP(&wikiName, "name", "n", "", "Space name")
	wikiSpaceUpdateCmd.Flags().StringVar(&wikiVisibility, "visibility", "", "Visibility")
	wikiSpaceUpdateCmd.Flags().StringVar(&wikiDesc, "desc", "", "Description")
	wikiSpaceUpdateCmd.MarkFlagRequired("id")

	wikiSpaceDeleteCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceDeleteCmd.MarkFlagRequired("id")

	// space member flags
	wikiSpaceMemberListCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceMemberListCmd.MarkFlagRequired("id")

	wikiSpaceMemberAddCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceMemberAddCmd.Flags().StringVar(&wikiMemberID, "member-id", "", "Member ID")
	wikiSpaceMemberAddCmd.Flags().StringVar(&wikiMemberType, "member-type", "", "Member type")
	wikiSpaceMemberAddCmd.Flags().StringVar(&wikiRoleID, "role-id", "", "Role ID")
	wikiSpaceMemberAddCmd.MarkFlagRequired("id")
	wikiSpaceMemberAddCmd.MarkFlagRequired("member-id")
	wikiSpaceMemberAddCmd.MarkFlagRequired("member-type")

	wikiSpaceMemberRemoveCmd.Flags().StringVarP(&wikiSpaceID, "id", "i", "", "Space ID")
	wikiSpaceMemberRemoveCmd.Flags().StringVar(&wikiMemberID, "member-id", "", "Member ID")
	wikiSpaceMemberRemoveCmd.MarkFlagRequired("id")
	wikiSpaceMemberRemoveCmd.MarkFlagRequired("member-id")

	// page flags
	wikiPageListCmd.Flags().StringVarP(&wikiSpaceIDFlag, "space-id", "s", "", "Space ID")
	wikiPageListCmd.MarkFlagRequired("space-id")

	wikiPageGetCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageGetCmd.MarkFlagRequired("id")

	wikiPageCreateCmd.Flags().StringVarP(&wikiSpaceIDFlag, "space-id", "s", "", "Space ID")
	wikiPageCreateCmd.Flags().StringVarP(&wikiTitle, "name", "n", "", "Page title")
	wikiPageCreateCmd.Flags().StringVar(&wikiParentID, "parent-id", "", "Parent page ID")
	wikiPageCreateCmd.Flags().StringVar(&wikiPageType, "type", "", "Page type (doc or folder)")
	wikiPageCreateCmd.MarkFlagRequired("space-id")
	wikiPageCreateCmd.MarkFlagRequired("name")

	wikiPageUpdateCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageUpdateCmd.Flags().StringVarP(&wikiTitle, "name", "n", "", "Page title")
	wikiPageUpdateCmd.MarkFlagRequired("id")

	wikiPageDeleteCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageDeleteCmd.MarkFlagRequired("id")

	// page content flags
	wikiPageContentGetCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageContentGetCmd.MarkFlagRequired("id")

	wikiPageContentUpdateCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageContentUpdateCmd.Flags().StringVar(&wikiContent, "content", "", "Page content")
	wikiPageContentUpdateCmd.Flags().StringVar(&wikiFormat, "format", "", "Content format (e.g. markdown)")
	wikiPageContentUpdateCmd.MarkFlagRequired("id")
	wikiPageContentUpdateCmd.MarkFlagRequired("content")

	// page versions flags
	wikiPageVersionsCmd.Flags().StringVarP(&wikiPageID, "id", "i", "", "Page ID")
	wikiPageVersionsCmd.MarkFlagRequired("id")
}
