package ship

import (
	"pingcode-client/internal/pkg/sdk"

	"github.com/spf13/cobra"
)

// Run executes the logic for ship
func Run(client *sdk.Client) {
	ListProducts(client)
}

// ── Product ────────────────────────────────────────────────────────────────

func ListProducts(client *sdk.Client) ([]sdk.Product, error) {
	return client.ListProducts()
}

func CreateProduct(client *sdk.Client, product *sdk.Product) (*sdk.Product, error) {
	return client.CreateProduct(product)
}

func UpdateProduct(client *sdk.Client, productID string, product *sdk.Product) (*sdk.Product, error) {
	return client.UpdateProduct(productID, product)
}

func ListProductMembers(client *sdk.Client, productID string) ([]sdk.ProductMember, error) {
	return client.ListProductMembers(productID)
}

func AddProductMember(client *sdk.Client, productID string, memberID string, memberType string, roleID string) (*sdk.ProductMember, error) {
	return client.AddProductMember(productID, memberID, memberType, roleID)
}

func RemoveProductMember(client *sdk.Client, productID string, memberID string) error {
	return client.RemoveProductMember(productID, memberID)
}

func AddProductSuite(client *sdk.Client, productID string, suite *sdk.IdeaSuite) (*sdk.IdeaSuite, error) {
	return client.AddProductSuite(productID, suite)
}

func ListProductSuites(client *sdk.Client, productID string) ([]sdk.IdeaSuite, error) {
	return client.ListProductSuites(productID)
}

func RemoveProductSuite(client *sdk.Client, productID string, suiteID string) error {
	return client.RemoveProductSuite(productID, suiteID)
}

func ListProductPlans(client *sdk.Client, productID string) ([]sdk.IdeaPlan, error) {
	return client.ListProductPlans(productID)
}

func ListProductChannels(client *sdk.Client, productID string) ([]sdk.ProductChannel, error) {
	return client.ListProductChannels(productID)
}

func ListProductTicketTypes(client *sdk.Client, productID string) ([]sdk.ProductTicketType, error) {
	return client.ListProductTicketTypes(productID)
}

func ListProductTags(client *sdk.Client, productID string) ([]sdk.ProductTag, error) {
	return client.ListProductTags(productID)
}

func AddProductTag(client *sdk.Client, productID string, name string) (*sdk.ProductTag, error) {
	return client.AddProductTag(productID, name)
}

func RemoveProductTag(client *sdk.Client, productID string, tagID string) error {
	return client.RemoveProductTag(productID, tagID)
}

// ── Idea ──────────────────────────────────────────────────────────────────

func ListIdeas(client *sdk.Client, productID string) ([]sdk.Idea, error) {
	return client.ListIdeas(productID)
}

func CreateIdea(client *sdk.Client, idea *sdk.Idea) (*sdk.Idea, error) {
	return client.CreateIdea(idea)
}

func UpdateIdea(client *sdk.Client, ideaID string, idea *sdk.Idea) (*sdk.Idea, error) {
	return client.UpdateIdea(ideaID, idea)
}

func ListIdeaStates(client *sdk.Client, productID string) ([]sdk.IdeaState, error) {
	return client.ListIdeaStates(productID)
}

func ListIdeaPriorities(client *sdk.Client, productID string) ([]sdk.IdeaPriority, error) {
	return client.ListIdeaPriorities(productID)
}

func ListIdeaProperties(client *sdk.Client, productID string) ([]sdk.IdeaProperty, error) {
	return client.ListIdeaProperties(productID)
}

func ListIdeaSuites(client *sdk.Client, productID string) ([]sdk.IdeaSuite, error) {
	return client.ListIdeaSuites(productID)
}

func ListIdeaPlans(client *sdk.Client, productID string) ([]sdk.IdeaPlan, error) {
	return client.ListIdeaPlans(productID)
}

func ListIdeaTransitionHistories(client *sdk.Client, ideaID string) ([]sdk.IdeaTransitionHistory, error) {
	return client.ListIdeaTransitionHistories(ideaID)
}

// ── Ticket ────────────────────────────────────────────────────────────────

func ListTickets(client *sdk.Client, productID, typeID, stateID, priorityID, keywords string) ([]sdk.Ticket, error) {
	return client.ListTickets(productID, typeID, stateID, priorityID, keywords)
}

func GetTicket(client *sdk.Client, ticketID string) (*sdk.Ticket, error) {
	return client.GetTicket(ticketID)
}

func CreateTicket(client *sdk.Client, productID, title, desc, typeID, assigneeID, priorityID, channelID, customerID string) (*sdk.Ticket, error) {
	return client.CreateTicket(&sdk.Ticket{
		ProductID:   productID,
		Title:       title,
		Description: desc,
		TypeID:      typeID,
		AssigneeID:  assigneeID,
		PriorityID:  priorityID,
		ChannelID:   channelID,
		CustomerID:  customerID,
	})
}

func UpdateTicket(client *sdk.Client, ticketID, title, desc, typeID, stateID, assigneeID, priorityID, solutionID, customerID string) (*sdk.Ticket, error) {
	return client.UpdateTicket(ticketID, &sdk.Ticket{
		Title:       title,
		Description: desc,
		TypeID:      typeID,
		StateID:     stateID,
		AssigneeID:  assigneeID,
		PriorityID:  priorityID,
		SolutionID:  solutionID,
		CustomerID:  customerID,
	})
}

func ListTicketPriorities(client *sdk.Client, productID string) ([]sdk.TicketPriority, error) {
	return client.ListTicketPriorities(productID)
}

func ListTicketStates(client *sdk.Client, productID string) ([]sdk.TicketState, error) {
	return client.ListTicketStates(productID)
}

func Help(cmd *cobra.Command, args []string) {
	cmd.Help()
}
