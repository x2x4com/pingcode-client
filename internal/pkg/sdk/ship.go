package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Identifier  string          `json:"identifier"`
	Description string          `json:"description"`
	Members     []ProductMember `json:"members,omitempty"`
	CreatedAt   int64           `json:"created_at,omitempty"`
	UpdatedAt   int64           `json:"updated_at,omitempty"`
}

type ProductMember struct {
	ID   string `json:"id"`
	Type string `json:"type"` // user, user_group
	User *struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"user,omitempty"`
	UserGroup *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user_group,omitempty"`
	Role *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"role,omitempty"`
}

type ProductTag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type ProductChannel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductTicketType struct {
	ID         string `json:"id"`
	TicketType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"ticket_type"`
}

type ProductListResponse struct {
	Values    []Product `json:"values"`
	Total     int       `json:"total"`
	PageSize  int       `json:"page_size"`
	PageIndex int       `json:"page_index"`
}

type ProductMemberListResponse struct {
	Values []ProductMember `json:"values"`
}

type ProductTagListResponse struct {
	Values []ProductTag `json:"values"`
}

type ProductChannelListResponse struct {
	Values []ProductChannel `json:"values"`
}

type ProductTicketTypeListResponse struct {
	Values []ProductTicketType `json:"values"`
}

type Idea struct {
	ID          string         `json:"id"`
	Identifier  string         `json:"identifier"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ProductID   string         `json:"product_id,omitempty"`
	AssigneeID  string         `json:"assignee_id,omitempty"`
	SuiteID     string         `json:"suite_id,omitempty"`
	PlanID      string         `json:"plan_id,omitempty"`
	StateID     string         `json:"state_id,omitempty"`
	PriorityID  string         `json:"priority_id,omitempty"`
	Progress    float64        `json:"progress"`
	State       *IdeaState     `json:"state,omitempty"`
	Priority    *IdeaPriority  `json:"priority,omitempty"`
	Suite       *IdeaSuite     `json:"suite,omitempty"`
	Plan        *IdeaPlan      `json:"plan,omitempty"`
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
	Properties  map[string]any `json:"properties,omitempty"`
}

type IdeaState struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type IdeaPriority struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IdeaSuite struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ParentID string `json:"parent_id,omitempty"`
}

type IdeaPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IdeaProperty struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Options []struct {
		ID   string `json:"_id"`
		Text string `json:"text"`
	} `json:"options"`
}

type IdeaTransitionHistory struct {
	ID        string     `json:"id"`
	FromState *IdeaState `json:"from_state"`
	ToState   *IdeaState `json:"to_state"`
	CreatedAt int64      `json:"created_at"`
}

type IdeaListResponse struct {
	Values    []Idea `json:"values"`
	Total     int    `json:"total"`
	PageSize  int    `json:"page_size"`
	PageIndex int    `json:"page_index"`
}

type IdeaStateListResponse struct {
	Values []IdeaState `json:"values"`
}

type IdeaPriorityListResponse struct {
	Values []IdeaPriority `json:"values"`
}

type IdeaSuiteListResponse struct {
	Values []IdeaSuite `json:"values"`
}

type IdeaPlanListResponse struct {
	Values []IdeaPlan `json:"values"`
}

type IdeaPropertyListResponse struct {
	Values []IdeaProperty `json:"values"`
}

type IdeaTransitionHistoryListResponse struct {
	Values []IdeaTransitionHistory `json:"values"`
}

func (c *Client) ListProducts() ([]Product, error) {
	resp, err := c.Request("GET", "/v1/ship/products", nil)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list products failed with status: %d", resp.StatusCode)
	}

	var listResp ProductListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) CreateProduct(product *Product) (*Product, error) {
	resp, err := c.Request("POST", "/v1/ship/products", product)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create product failed with status: %d", resp.StatusCode)
	}

	var newProduct Product
	if err := json.NewDecoder(resp.Body).Decode(&newProduct); err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func (c *Client) UpdateProduct(productID string, product *Product) (*Product, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/v1/ship/products/%s", productID), product)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update product failed with status: %d", resp.StatusCode)
	}

	var updatedProduct Product
	if err := json.NewDecoder(resp.Body).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func (c *Client) ListProductMembers(productID string) ([]ProductMember, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/members", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product members failed with status: %d", resp.StatusCode)
	}

	var listResp ProductMemberListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) AddProductMember(productID string, memberID string, memberType string, roleID string) (*ProductMember, error) {
	payload := map[string]any{
		"member": map[string]string{
			"id":   memberID,
			"type": memberType,
		},
	}
	if roleID != "" {
		payload["role_id"] = roleID
	}

	resp, err := c.Request("POST", fmt.Sprintf("/v1/ship/products/%s/members", productID), payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("add product member failed with status: %d", resp.StatusCode)
	}

	var newMember ProductMember
	if err := json.NewDecoder(resp.Body).Decode(&newMember); err != nil {
		return nil, err
	}

	return &newMember, nil
}

func (c *Client) RemoveProductMember(productID string, memberID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/v1/ship/products/%s/members/%s", productID, memberID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("remove product member failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) AddProductSuite(productID string, suite *IdeaSuite) (*IdeaSuite, error) {
	resp, err := c.Request("POST", fmt.Sprintf("/v1/ship/products/%s/suites", productID), suite)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("add product suite failed with status: %d", resp.StatusCode)
	}

	var newSuite IdeaSuite
	if err := json.NewDecoder(resp.Body).Decode(&newSuite); err != nil {
		return nil, err
	}

	return &newSuite, nil
}

func (c *Client) RemoveProductSuite(productID string, suiteID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/v1/ship/products/%s/suites/%s", productID, suiteID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("remove product suite failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListProductSuites(productID string) ([]IdeaSuite, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/suites", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product suites failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaSuiteListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListProductPlans(productID string) ([]IdeaPlan, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/plans", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product plans failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaPlanListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListProductChannels(productID string) ([]ProductChannel, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/channels", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product channels failed with status: %d", resp.StatusCode)
	}

	var listResp ProductChannelListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListProductTicketTypes(productID string) ([]ProductTicketType, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/ticket_types", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product ticket types failed with status: %d", resp.StatusCode)
	}

	var listResp ProductTicketTypeListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListProductTags(productID string) ([]ProductTag, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/products/%s/tags", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list product tags failed with status: %d", resp.StatusCode)
	}

	var listResp ProductTagListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) AddProductTag(productID string, name string) (*ProductTag, error) {
	payload := map[string]string{
		"name": name,
	}
	resp, err := c.Request("POST", fmt.Sprintf("/v1/ship/products/%s/tags", productID), payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("add product tag failed with status: %d", resp.StatusCode)
	}

	var newTag ProductTag
	if err := json.NewDecoder(resp.Body).Decode(&newTag); err != nil {
		return nil, err
	}

	return &newTag, nil
}

func (c *Client) RemoveProductTag(productID string, tagID string) error {
	resp, err := c.Request("DELETE", fmt.Sprintf("/v1/ship/products/%s/tags/%s", productID, tagID), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("remove product tag failed with status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ListIdeas(productID string) ([]Idea, error) {
	url := "/v1/ship/ideas"
	if productID != "" {
		url += "?product_id=" + productID
	}
	resp, err := c.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list ideas failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) CreateIdea(idea *Idea) (*Idea, error) {
	resp, err := c.Request("POST", "/v1/ship/ideas", idea)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create idea failed with status: %d", resp.StatusCode)
	}

	var newIdea Idea
	if err := json.NewDecoder(resp.Body).Decode(&newIdea); err != nil {
		return nil, err
	}

	return &newIdea, nil
}

func (c *Client) UpdateIdea(ideaID string, idea *Idea) (*Idea, error) {
	resp, err := c.Request("PATCH", fmt.Sprintf("/v1/ship/ideas/%s", ideaID), idea)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update idea failed with status: %d", resp.StatusCode)
	}

	var updatedIdea Idea
	if err := json.NewDecoder(resp.Body).Decode(&updatedIdea); err != nil {
		return nil, err
	}

	return &updatedIdea, nil
}

func (c *Client) ListIdeaStates(productID string) ([]IdeaState, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/idea/states?product_id=%s", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea states failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaStateListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListIdeaPriorities(productID string) ([]IdeaPriority, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/idea/priorities?product_id=%s", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea priorities failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaPriorityListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListIdeaProperties(productID string) ([]IdeaProperty, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/idea/properties?product_id=%s", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea properties failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaPropertyListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListIdeaSuites(productID string) ([]IdeaSuite, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/idea/suites?product_id=%s", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea suites failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaSuiteListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListIdeaPlans(productID string) ([]IdeaPlan, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/idea/plans?product_id=%s", productID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea plans failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaPlanListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}

func (c *Client) ListIdeaTransitionHistories(ideaID string) ([]IdeaTransitionHistory, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v1/ship/ideas/%s/transition_histories", ideaID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list idea transition histories failed with status: %d", resp.StatusCode)
	}

	var listResp IdeaTransitionHistoryListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}
