package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"pingcode-client/internal/app/ship"
	"pingcode-client/internal/pkg/output"
	"pingcode-client/internal/pkg/sdk"

	"github.com/spf13/cobra"
)

var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "PingCode 产品管理 (Ship) 模块",
	Long:  `用于管理和列出 PingCode 中的产品信息，包含产品的名称、ID 以及描述等。`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			panic(err)
		}
		ship.Run(client)
	},
}

var (
	productID  string
	ideaID     string
	stateID    string
	priorityID string
	assigneeID string
	suiteID    string
	planID     string

	// Product specific flags
	pName        string
	pIdentifier  string
	pDescription string
	pMemberID    string
	pMemberType  string
	pRoleID      string
	pTagName     string
	pTagID       string
	pParentID    string
	pSuiteType   string

	// Ticket specific flags
	ticketID   string
	solutionID string
	customerID string
	channelID  string
	keywords   string
)

// toMap converts a struct to map[string]interface{} using JSON serialization
func toMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}

// toMapSlice converts a slice of structs to []map[string]interface{}
func toMapSlice(v interface{}) []map[string]interface{} {
	b, _ := json.Marshal(v)
	var m []map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}

var productCmd = &cobra.Command{
	Use:   "product",
	Short: "产品管理",
}

var listProductsCmd = &cobra.Command{
	Use:   "list",
	Short: "获取产品列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProducts(client)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var createProductCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个产品",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.CreateProduct(client, &sdk.Product{
			Name:        pName,
			Identifier:  pIdentifier,
			Description: pDescription,
		})
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var updateProductCmd = &cobra.Command{
	Use:   "update",
	Short: "部分更新一个产品",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.UpdateProduct(client, productID, &sdk.Product{
			Name:        pName,
			Identifier:  pIdentifier,
			Description: pDescription,
		})
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var productMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "产品成员管理",
}

var listProductMembersCmd = &cobra.Command{
	Use:   "list",
	Short: "获取产品中的成员列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductMembers(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var addProductMemberCmd = &cobra.Command{
	Use:   "add",
	Short: "向产品中添加一个成员",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.AddProductMember(client, productID, pMemberID, pMemberType, pRoleID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var removeProductMemberCmd = &cobra.Command{
	Use:   "remove",
	Short: "在产品中移除一个成员",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		err = ship.RemoveProductMember(client, productID, pMemberID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Member removed successfully")
	},
}

var productModulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "产品需求模块管理",
}

var listProductModulesCmd = &cobra.Command{
	Use:   "list",
	Short: "获取产品中的需求模块列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductSuites(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var addProductModuleCmd = &cobra.Command{
	Use:   "add",
	Short: "向产品中添加一个需求模块",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.AddProductSuite(client, productID, &sdk.IdeaSuite{
			Name:     pName,
			Type:     pSuiteType,
			ParentID: pParentID,
		})
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var removeProductModuleCmd = &cobra.Command{
	Use:   "remove",
	Short: "在产品中移除一个需求模块",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		err = ship.RemoveProductSuite(client, productID, suiteID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Suite removed successfully")
	},
}

var listProductPlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "获取产品中的需求排期列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductPlans(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listProductChannelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "获取产品中的工单渠道列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductChannels(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listProductTicketTypesCmd = &cobra.Command{
	Use:   "ticket-types",
	Short: "获取产品中的工单类型列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductTicketTypes(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var productTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "产品标签管理",
}

var listProductTagsCmd = &cobra.Command{
	Use:   "list",
	Short: "获取产品中的标签列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListProductTags(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var addProductTagCmd = &cobra.Command{
	Use:   "add",
	Short: "向产品中添加一个标签",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.AddProductTag(client, productID, pTagName)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var removeProductTagCmd = &cobra.Command{
	Use:   "remove",
	Short: "在产品中移除一个标签",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		err = ship.RemoveProductTag(client, productID, pTagID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tag removed successfully")
	},
}

var ideasCmd = &cobra.Command{
	Use:   "ideas",
	Short: "需求管理",
}

var listIdeasCmd = &cobra.Command{
	Use:   "list",
	Short: "获取需求列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeas(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var createIdeaCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个需求",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.CreateIdea(client, &sdk.Idea{
			ProductID:   productID,
			Title:       title,
			Description: desc,
			AssigneeID:  assigneeID,
			SuiteID:     suiteID,
			PriorityID:  priorityID,
		})
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var updateIdeaCmd = &cobra.Command{
	Use:   "update",
	Short: "部分更新一个需求",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.UpdateIdea(client, ideaID, &sdk.Idea{
			Title:       title,
			Description: desc,
			StateID:     stateID,
			PriorityID:  priorityID,
			AssigneeID:  assigneeID,
			SuiteID:     suiteID,
			PlanID:      planID,
		})
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listStatesCmd = &cobra.Command{
	Use:   "states",
	Short: "获取需求状态列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaStates(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listPrioritiesCmd = &cobra.Command{
	Use:   "priorities",
	Short: "获取需求优先级列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaPriorities(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listPropertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "获取需求属性列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaProperties(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listSuitesCmd = &cobra.Command{
	Use:   "modules",
	Short: "获取需求模块列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaSuites(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listPlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "获取需求排期列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaPlans(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listHistoriesCmd = &cobra.Command{
	Use:   "histories",
	Short: "获取需求流转记录列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListIdeaTransitionHistories(client, ideaID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var ticketsCmd = &cobra.Command{Use: "tickets", Short: "工单管理"}

var listTicketsCmd = &cobra.Command{
	Use:   "list",
	Short: "获取工单列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListTickets(client, productID, typeID, stateID, priorityID, keywords)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var getTicketCmd = &cobra.Command{
	Use:   "get",
	Short: "获取工单详情",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.GetTicket(client, ticketID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var createTicketCmd = &cobra.Command{
	Use:   "create",
	Short: "创建工单",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.CreateTicket(client, productID, title, desc, typeID, assigneeID, priorityID, channelID, customerID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var updateTicketCmd = &cobra.Command{
	Use:   "update",
	Short: "更新工单",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.UpdateTicket(client, ticketID, title, desc, typeID, stateID, assigneeID, priorityID, solutionID, customerID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := output.FormatAndPrint(toMap(data), opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listTicketPrioritiesCmd = &cobra.Command{
	Use:   "priorities",
	Short: "获取工单优先级列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListTicketPriorities(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var listTicketStatesCmd = &cobra.Command{
	Use:   "states",
	Short: "获取工单状态列表",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ship.ListTicketStates(client, productID)
		if err != nil {
			log.Fatal(err)
		}
		opts := GetOutputOptions()
		if opts.Raw {
			if err := output.FormatAndPrint(data, opts); err != nil {
				log.Fatal(err)
			}
		} else {
			simplified := toMapSlice(data)
			if err := output.FormatAndPrint(simplified, opts); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(shipCmd)

	// Product commands
	shipCmd.AddCommand(productCmd)
	productCmd.AddCommand(listProductsCmd)
	productCmd.AddCommand(createProductCmd)
	productCmd.AddCommand(updateProductCmd)

	productCmd.AddCommand(productMembersCmd)
	productMembersCmd.AddCommand(listProductMembersCmd)
	productMembersCmd.AddCommand(addProductMemberCmd)
	productMembersCmd.AddCommand(removeProductMemberCmd)

	productCmd.AddCommand(productModulesCmd)
	productModulesCmd.AddCommand(listProductModulesCmd)
	productModulesCmd.AddCommand(addProductModuleCmd)
	productModulesCmd.AddCommand(removeProductModuleCmd)

	productCmd.AddCommand(listProductPlansCmd)
	productCmd.AddCommand(listProductChannelsCmd)
	productCmd.AddCommand(listProductTicketTypesCmd)

	productCmd.AddCommand(productTagsCmd)
	productTagsCmd.AddCommand(listProductTagsCmd)
	productTagsCmd.AddCommand(addProductTagCmd)
	productTagsCmd.AddCommand(removeProductTagCmd)

	// Product flags
	productCmd.PersistentFlags().StringVar(&productID, "product-id", "", "产品 ID")
	createProductCmd.Flags().StringVar(&pName, "name", "", "产品名称")
	createProductCmd.Flags().StringVar(&pIdentifier, "identifier", "", "产品标识")
	createProductCmd.Flags().StringVar(&pDescription, "desc", "", "产品描述")

	updateProductCmd.Flags().StringVar(&pName, "name", "", "产品名称")
	updateProductCmd.Flags().StringVar(&pIdentifier, "identifier", "", "产品标识")
	updateProductCmd.Flags().StringVar(&pDescription, "desc", "", "产品描述")

	addProductMemberCmd.Flags().StringVar(&pMemberID, "member-id", "", "成员 ID")
	addProductMemberCmd.Flags().StringVar(&pMemberType, "member-type", "user", "成员类型 (user/user_group)")
	addProductMemberCmd.Flags().StringVar(&pRoleID, "role-id", "", "角色 ID")
	removeProductMemberCmd.Flags().StringVar(&pMemberID, "member-id", "", "成员 ID")

	addProductModuleCmd.Flags().StringVar(&pName, "name", "", "模块名称")
	addProductModuleCmd.Flags().StringVar(&pSuiteType, "type", "module", "模块类型 (module/product)")
	addProductModuleCmd.Flags().StringVar(&pParentID, "parent-id", "", "父模块 ID")
	removeProductModuleCmd.Flags().StringVar(&suiteID, "suite-id", "", "模块 ID")

	addProductTagCmd.Flags().StringVar(&pTagName, "name", "", "标签名称")
	removeProductTagCmd.Flags().StringVar(&pTagID, "tag-id", "", "标签 ID")

	// Idea commands
	shipCmd.AddCommand(ideasCmd)
	ideasCmd.AddCommand(listIdeasCmd)
	ideasCmd.AddCommand(createIdeaCmd)
	ideasCmd.AddCommand(updateIdeaCmd)
	ideasCmd.AddCommand(listStatesCmd)
	ideasCmd.AddCommand(listPrioritiesCmd)
	ideasCmd.AddCommand(listPropertiesCmd)
	ideasCmd.AddCommand(listSuitesCmd)
	ideasCmd.AddCommand(listPlansCmd)
	ideasCmd.AddCommand(listHistoriesCmd)

	// Idea flags
	ideasCmd.PersistentFlags().StringVar(&productID, "product-id", "", "产品 ID")
	ideasCmd.PersistentFlags().StringVar(&ideaID, "idea-id", "", "需求 ID")

	createIdeaCmd.Flags().StringVar(&title, "title", "", "需求标题")
	createIdeaCmd.Flags().StringVar(&desc, "desc", "", "需求描述")
	createIdeaCmd.Flags().StringVar(&assigneeID, "assignee-id", "", "负责人 ID")
	createIdeaCmd.Flags().StringVar(&suiteID, "suite-id", "", "模块 ID")
	createIdeaCmd.Flags().StringVar(&priorityID, "priority-id", "", "优先级 ID")

	updateIdeaCmd.Flags().StringVar(&title, "title", "", "需求标题")
	updateIdeaCmd.Flags().StringVar(&desc, "desc", "", "需求描述")
	updateIdeaCmd.Flags().StringVar(&stateID, "state-id", "", "状态 ID")
	updateIdeaCmd.Flags().StringVar(&priorityID, "priority-id", "", "优先级 ID")
	updateIdeaCmd.Flags().StringVar(&assigneeID, "assignee-id", "", "负责人 ID")
	updateIdeaCmd.Flags().StringVar(&suiteID, "suite-id", "", "模块 ID")
	updateIdeaCmd.Flags().StringVar(&planID, "plan-id", "", "排期 ID")

	// Ticket commands
	shipCmd.AddCommand(ticketsCmd)
	ticketsCmd.AddCommand(listTicketsCmd)
	ticketsCmd.AddCommand(getTicketCmd)
	ticketsCmd.AddCommand(createTicketCmd)
	ticketsCmd.AddCommand(updateTicketCmd)
	ticketsCmd.AddCommand(listTicketPrioritiesCmd)
	ticketsCmd.AddCommand(listTicketStatesCmd)

	ticketsCmd.PersistentFlags().StringVar(&productID, "product-id", "", "产品 ID")

	listTicketsCmd.Flags().StringVar(&typeID, "type-id", "", "工单类型 ID")
	listTicketsCmd.Flags().StringVar(&stateID, "state-id", "", "工单状态 ID")
	listTicketsCmd.Flags().StringVar(&priorityID, "priority-id", "", "工单优先级 ID")
	listTicketsCmd.Flags().StringVar(&keywords, "keywords", "", "关键字搜索")

	getTicketCmd.Flags().StringVar(&ticketID, "id", "", "工单 ID")
	getTicketCmd.MarkFlagRequired("id")

	createTicketCmd.Flags().StringVar(&title, "title", "", "工单标题")
	createTicketCmd.Flags().StringVar(&desc, "desc", "", "工单描述")
	createTicketCmd.Flags().StringVar(&typeID, "type-id", "", "工单类型 ID")
	createTicketCmd.Flags().StringVar(&assigneeID, "assignee-id", "", "负责人 ID")
	createTicketCmd.Flags().StringVar(&priorityID, "priority-id", "", "优先级 ID")
	createTicketCmd.Flags().StringVar(&channelID, "channel-id", "", "渠道 ID")
	createTicketCmd.Flags().StringVar(&customerID, "customer-id", "", "客户 ID")
	createTicketCmd.MarkFlagRequired("title")
	createTicketCmd.MarkFlagRequired("type-id")

	updateTicketCmd.Flags().StringVar(&ticketID, "id", "", "工单 ID")
	updateTicketCmd.Flags().StringVar(&title, "title", "", "工单标题")
	updateTicketCmd.Flags().StringVar(&desc, "desc", "", "工单描述")
	updateTicketCmd.Flags().StringVar(&typeID, "type-id", "", "工单类型 ID")
	updateTicketCmd.Flags().StringVar(&stateID, "state-id", "", "工单状态 ID")
	updateTicketCmd.Flags().StringVar(&assigneeID, "assignee-id", "", "负责人 ID")
	updateTicketCmd.Flags().StringVar(&priorityID, "priority-id", "", "工单优先级 ID")
	updateTicketCmd.Flags().StringVar(&solutionID, "solution-id", "", "解决方案 ID")
	updateTicketCmd.Flags().StringVar(&customerID, "customer-id", "", "客户 ID")
	updateTicketCmd.MarkFlagRequired("id")
}
