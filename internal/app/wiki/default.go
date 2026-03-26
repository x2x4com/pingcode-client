package wiki

import (
	"fmt"
	"log"
	"pingcode-client/internal/pkg/sdk"
)

func ListSpaces(c *sdk.Client) {
	spaces, err := c.ListWikiSpaces(0, 30)
	if err != nil {
		log.Fatalf("Error listing wiki spaces: %v", err)
	}
	fmt.Println("Wiki Spaces:")
	for _, s := range spaces {
		fmt.Printf("- [%s] %s (identifier: %s, scope_type: %s)\n", s.ID, s.Name, s.Identifier, s.ScopeType)
	}
}

func GetSpace(c *sdk.Client, spaceID string) {
	s, err := c.GetWikiSpace(spaceID)
	if err != nil {
		log.Fatalf("Error getting wiki space: %v", err)
	}
	fmt.Printf("ID:          %s\nName:        %s\nIdentifier:  %s\nScope Type:  %s\nScope ID:    %s\nVisibility:  %s\nDescription: %s\n",
		s.ID, s.Name, s.Identifier, s.ScopeType, s.ScopeID, s.Visibility, s.Description)
}

func CreateSpace(c *sdk.Client, name, identifier, scopeType, scopeID, visibility, desc string) {
	s, err := c.CreateWikiSpace(name, identifier, scopeType, scopeID, visibility, desc)
	if err != nil {
		log.Fatalf("Error creating wiki space: %v", err)
	}
	fmt.Printf("Created Wiki Space: %s (ID: %s)\n", s.Name, s.ID)
}

func UpdateSpace(c *sdk.Client, spaceID, name, visibility, desc string) {
	s, err := c.UpdateWikiSpace(spaceID, name, visibility, desc)
	if err != nil {
		log.Fatalf("Error updating wiki space: %v", err)
	}
	fmt.Printf("Updated Wiki Space: %s (ID: %s)\n", s.Name, s.ID)
}

func DeleteSpace(c *sdk.Client, spaceID string) {
	if err := c.DeleteWikiSpace(spaceID); err != nil {
		log.Fatalf("Error deleting wiki space: %v", err)
	}
	fmt.Printf("Deleted Wiki Space: %s\n", spaceID)
}

func ListSpaceMembers(c *sdk.Client, spaceID string) {
	members, err := c.ListWikiSpaceMembers(spaceID)
	if err != nil {
		log.Fatalf("Error listing wiki space members: %v", err)
	}
	fmt.Printf("Wiki Space Members (space: %s):\n", spaceID)
	for _, m := range members {
		fmt.Printf("- [%s] type: %s\n", m.ID, m.Type)
	}
}

func AddSpaceMember(c *sdk.Client, spaceID, memberID, memberType, roleID string) {
	m, err := c.AddWikiSpaceMember(spaceID, memberID, memberType, roleID)
	if err != nil {
		log.Fatalf("Error adding wiki space member: %v", err)
	}
	fmt.Printf("Added Member: %s (type: %s) to Space: %s\n", m.ID, m.Type, spaceID)
}

func RemoveSpaceMember(c *sdk.Client, spaceID, memberID string) {
	if err := c.RemoveWikiSpaceMember(spaceID, memberID); err != nil {
		log.Fatalf("Error removing wiki space member: %v", err)
	}
	fmt.Printf("Removed Member: %s from Space: %s\n", memberID, spaceID)
}

func ListPages(c *sdk.Client, spaceID string) {
	pages, err := c.ListWikiPages(spaceID, 0, 30)
	if err != nil {
		log.Fatalf("Error listing wiki pages: %v", err)
	}
	fmt.Printf("Wiki Pages (space: %s):\n", spaceID)
	for _, p := range pages {
		fmt.Printf("- [%s] %s (type: %s)\n", p.ID, p.Title, p.Type)
	}
}

func GetPage(c *sdk.Client, pageID string) {
	p, err := c.GetWikiPage(pageID)
	if err != nil {
		log.Fatalf("Error getting wiki page: %v", err)
	}
	fmt.Printf("ID:       %s\nTitle:    %s\nType:     %s\nSpaceID:  %s\nParentID: %s\n",
		p.ID, p.Title, p.Type, p.SpaceID, p.ParentID)
}

func CreatePage(c *sdk.Client, spaceID, parentID, title, pageType string) {
	p, err := c.CreateWikiPage(spaceID, parentID, title, pageType)
	if err != nil {
		log.Fatalf("Error creating wiki page: %v", err)
	}
	fmt.Printf("Created Wiki Page: %s (ID: %s)\n", p.Title, p.ID)
}

func UpdatePage(c *sdk.Client, pageID, title string) {
	p, err := c.UpdateWikiPage(pageID, title)
	if err != nil {
		log.Fatalf("Error updating wiki page: %v", err)
	}
	fmt.Printf("Updated Wiki Page: %s (ID: %s)\n", p.Title, p.ID)
}

func DeletePage(c *sdk.Client, pageID string) {
	if err := c.DeleteWikiPage(pageID); err != nil {
		log.Fatalf("Error deleting wiki page: %v", err)
	}
	fmt.Printf("Deleted Wiki Page: %s\n", pageID)
}

func GetPageContent(c *sdk.Client, pageID string) {
	content, err := c.GetWikiPageContent(pageID)
	if err != nil {
		log.Fatalf("Error getting wiki page content: %v", err)
	}
	fmt.Printf("PageID: %s\nFormat: %s\nContent:\n%s\n", content.PageID, content.Format, content.Content)
}

func UpdatePageContent(c *sdk.Client, pageID, content, format string) {
	c2, err := c.UpdateWikiPageContent(pageID, content, format)
	if err != nil {
		log.Fatalf("Error updating wiki page content: %v", err)
	}
	fmt.Printf("Updated content for Page: %s\n", c2.PageID)
}

func ListPageVersions(c *sdk.Client, pageID string) {
	versions, err := c.ListWikiPageVersions(pageID)
	if err != nil {
		log.Fatalf("Error listing wiki page versions: %v", err)
	}
	fmt.Printf("Wiki Page Versions (page: %s):\n", pageID)
	for _, v := range versions {
		fmt.Printf("- [%s] version: %d, created_at: %d\n", v.ID, v.Version, v.CreatedAt)
	}
}
