package ship

import (
	"fmt"
	"log"
	"pingcode-client/internal/pkg/sdk"

	"github.com/spf13/cobra"
)

// Run executes the logic for ship
func Run(client *sdk.Client) {
	ListProducts(client)
}

func ListProducts(client *sdk.Client) {
	log.Println("Listing products from Ship...")
	products, err := client.ListProducts()
	if err != nil {
		log.Fatalf("Error listing products: %v", err)
	}

	fmt.Println("Products:")
	for _, p := range products {
		fmt.Printf("- %s (ID: %s, Identifier: %s)\n", p.Name, p.ID, p.Identifier)
	}
}

func CreateProduct(client *sdk.Client, product *sdk.Product) {
	newProduct, err := client.CreateProduct(product)
	if err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
	fmt.Printf("Created Product: %s (ID: %s, Identifier: %s)\n", newProduct.Name, newProduct.ID, newProduct.Identifier)
}

func UpdateProduct(client *sdk.Client, productID string, product *sdk.Product) {
	updatedProduct, err := client.UpdateProduct(productID, product)
	if err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
	fmt.Printf("Updated Product: %s (ID: %s)\n", updatedProduct.Name, updatedProduct.ID)
}

func ListProductMembers(client *sdk.Client, productID string) {
	members, err := client.ListProductMembers(productID)
	if err != nil {
		log.Fatalf("Error listing product members: %v", err)
	}
	fmt.Println("Product Members:")
	for _, m := range members {
		name := ""
		if m.User != nil {
			name = m.User.DisplayName
		} else if m.UserGroup != nil {
			name = m.UserGroup.Name
		}
		roleName := ""
		if m.Role != nil {
			roleName = m.Role.Name
		}
		fmt.Printf("- %s (ID: %s, Type: %s, Role: %s)\n", name, m.ID, m.Type, roleName)
	}
}

func AddProductMember(client *sdk.Client, productID string, memberID string, memberType string, roleID string) {
	member, err := client.AddProductMember(productID, memberID, memberType, roleID)
	if err != nil {
		log.Fatalf("Error adding product member: %v", err)
	}
	fmt.Printf("Added Member: %s (ID: %s) to Product: %s\n", member.ID, member.Type, productID)
}

func RemoveProductMember(client *sdk.Client, productID string, memberID string) {
	err := client.RemoveProductMember(productID, memberID)
	if err != nil {
		log.Fatalf("Error removing product member: %v", err)
	}
	fmt.Printf("Removed Member: %s from Product: %s\n", memberID, productID)
}

func AddProductSuite(client *sdk.Client, productID string, suite *sdk.IdeaSuite) {
	newSuite, err := client.AddProductSuite(productID, suite)
	if err != nil {
		log.Fatalf("Error adding product suite: %v", err)
	}
	fmt.Printf("Added Suite: %s (ID: %s) to Product: %s\n", newSuite.Name, newSuite.ID, productID)
}

func ListProductSuites(client *sdk.Client, productID string) {
	suites, err := client.ListProductSuites(productID)
	if err != nil {
		log.Fatalf("Error listing product suites: %v", err)
	}
	fmt.Println("Product Modules (Suites):")
	for _, s := range suites {
		fmt.Printf("- %s (ID: %s, Type: %s)\n", s.Name, s.ID, s.Type)
	}
}

func RemoveProductSuite(client *sdk.Client, productID string, suiteID string) {
	err := client.RemoveProductSuite(productID, suiteID)
	if err != nil {
		log.Fatalf("Error removing product suite: %v", err)
	}
	fmt.Printf("Removed Suite: %s from Product: %s\n", suiteID, productID)
}

func ListProductPlans(client *sdk.Client, productID string) {
	plans, err := client.ListProductPlans(productID)
	if err != nil {
		log.Fatalf("Error listing product plans: %v", err)
	}
	fmt.Println("Product Plans:")
	for _, p := range plans {
		fmt.Printf("- %s (ID: %s)\n", p.Name, p.ID)
	}
}

func ListProductChannels(client *sdk.Client, productID string) {
	channels, err := client.ListProductChannels(productID)
	if err != nil {
		log.Fatalf("Error listing product channels: %v", err)
	}
	fmt.Println("Product Channels:")
	for _, c := range channels {
		fmt.Printf("- %s (ID: %s)\n", c.Name, c.ID)
	}
}

func ListProductTicketTypes(client *sdk.Client, productID string) {
	types, err := client.ListProductTicketTypes(productID)
	if err != nil {
		log.Fatalf("Error listing product ticket types: %v", err)
	}
	fmt.Println("Product Ticket Types:")
	for _, t := range types {
		fmt.Printf("- %s (ID: %s)\n", t.TicketType.Name, t.ID)
	}
}

func ListProductTags(client *sdk.Client, productID string) {
	tags, err := client.ListProductTags(productID)
	if err != nil {
		log.Fatalf("Error listing product tags: %v", err)
	}
	fmt.Println("Product Tags:")
	for _, t := range tags {
		fmt.Printf("- %s (ID: %s, Color: %s)\n", t.Name, t.ID, t.Color)
	}
}

func AddProductTag(client *sdk.Client, productID string, name string) {
	tag, err := client.AddProductTag(productID, name)
	if err != nil {
		log.Fatalf("Error adding product tag: %v", err)
	}
	fmt.Printf("Added Tag: %s (ID: %s) to Product: %s\n", tag.Name, tag.ID, productID)
}

func RemoveProductTag(client *sdk.Client, productID string, tagID string) {
	err := client.RemoveProductTag(productID, tagID)
	if err != nil {
		log.Fatalf("Error removing product tag: %v", err)
	}
	fmt.Printf("Removed Tag: %s from Product: %s\n", tagID, productID)
}

func ListIdeas(client *sdk.Client, productID string) {
	ideas, err := client.ListIdeas(productID)
	if err != nil {
		log.Fatalf("Error listing ideas: %v", err)
	}

	fmt.Println("Ideas:")
	for _, i := range ideas {
		stateName := ""
		if i.State != nil {
			stateName = i.State.Name
		}
		fmt.Printf("- [%s] %s (ID: %s, State: %s)\n", i.Identifier, i.Title, i.ID, stateName)
	}
}

func CreateIdea(client *sdk.Client, idea *sdk.Idea) {
	newIdea, err := client.CreateIdea(idea)
	if err != nil {
		log.Fatalf("Error creating idea: %v", err)
	}
	fmt.Printf("Created Idea: %s (ID: %s, Identifier: %s)\n", newIdea.Title, newIdea.ID, newIdea.Identifier)
}

func UpdateIdea(client *sdk.Client, ideaID string, idea *sdk.Idea) {
	updatedIdea, err := client.UpdateIdea(ideaID, idea)
	if err != nil {
		log.Fatalf("Error updating idea: %v", err)
	}
	fmt.Printf("Updated Idea: %s (ID: %s)\n", updatedIdea.Title, updatedIdea.ID)
}

func ListIdeaStates(client *sdk.Client, productID string) {
	states, err := client.ListIdeaStates(productID)
	if err != nil {
		log.Fatalf("Error listing idea states: %v", err)
	}
	fmt.Println("Idea States:")
	for _, s := range states {
		fmt.Printf("- %s (ID: %s, Type: %s)\n", s.Name, s.ID, s.Type)
	}
}

func ListIdeaPriorities(client *sdk.Client, productID string) {
	priorities, err := client.ListIdeaPriorities(productID)
	if err != nil {
		log.Fatalf("Error listing idea priorities: %v", err)
	}
	fmt.Println("Idea Priorities:")
	for _, p := range priorities {
		fmt.Printf("- %s (ID: %s)\n", p.Name, p.ID)
	}
}

func ListIdeaProperties(client *sdk.Client, productID string) {
	properties, err := client.ListIdeaProperties(productID)
	if err != nil {
		log.Fatalf("Error listing idea properties: %v", err)
	}
	fmt.Println("Idea Properties:")
	for _, p := range properties {
		fmt.Printf("- %s (ID: %s, Type: %s)\n", p.Name, p.ID, p.Type)
	}
}

func ListIdeaSuites(client *sdk.Client, productID string) {
	suites, err := client.ListIdeaSuites(productID)
	if err != nil {
		log.Fatalf("Error listing idea suites: %v", err)
	}
	fmt.Println("Idea Modules (Suites):")
	for _, s := range suites {
		fmt.Printf("- %s (ID: %s, Type: %s)\n", s.Name, s.ID, s.Type)
	}
}

func ListIdeaPlans(client *sdk.Client, productID string) {
	plans, err := client.ListIdeaPlans(productID)
	if err != nil {
		log.Fatalf("Error listing idea plans: %v", err)
	}
	fmt.Println("Idea Plans:")
	for _, p := range plans {
		fmt.Printf("- %s (ID: %s)\n", p.Name, p.ID)
	}
}

func ListIdeaTransitionHistories(client *sdk.Client, ideaID string) {
	histories, err := client.ListIdeaTransitionHistories(ideaID)
	if err != nil {
		log.Fatalf("Error listing idea transition histories: %v", err)
	}
	fmt.Println("Idea Transition Histories:")
	for _, h := range histories {
		fromName := "None"
		if h.FromState != nil {
			fromName = h.FromState.Name
		}
		toName := "None"
		if h.ToState != nil {
			toName = h.ToState.Name
		}
		fmt.Printf("- From: %s -> To: %s (ID: %s)\n", fromName, toName, h.ID)
	}
}

func Help(cmd *cobra.Command, args []string) {
	// ctx := cmd.Context()
	cmd.Help()
}
