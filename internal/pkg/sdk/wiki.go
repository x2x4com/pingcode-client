package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WikiSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type WikiSpaceListResponse struct {
	Values    []WikiSpace `json:"values"`
	Total     int         `json:"total"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
}

func (c *Client) ListWikiSpaces(pageIndex, pageSize int) ([]WikiSpace, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/wiki/spaces?page_index=%d&page_size=%d", pageIndex, pageSize), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list wiki spaces failed with status: %d", resp.StatusCode)
	}

	var listResp WikiSpaceListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return listResp.Values, nil
}
