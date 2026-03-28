package wiki

import (
	"pingcode-client/internal/pkg/sdk"
)

func ListSpaces(c *sdk.Client) ([]sdk.WikiSpace, error) {
	return c.ListWikiSpaces(0, 30)
}

func GetSpace(c *sdk.Client, spaceID string) (*sdk.WikiSpace, error) {
	return c.GetWikiSpace(spaceID)
}

func CreateSpace(c *sdk.Client, name, identifier, scopeType, scopeID, visibility, desc string) (*sdk.WikiSpace, error) {
	return c.CreateWikiSpace(name, identifier, scopeType, scopeID, visibility, desc)
}

func UpdateSpace(c *sdk.Client, spaceID, name, visibility, desc string) (*sdk.WikiSpace, error) {
	return c.UpdateWikiSpace(spaceID, name, visibility, desc)
}

func DeleteSpace(c *sdk.Client, spaceID string) error {
	return c.DeleteWikiSpace(spaceID)
}

func ListSpaceMembers(c *sdk.Client, spaceID string) ([]sdk.WikiSpaceMember, error) {
	return c.ListWikiSpaceMembers(spaceID)
}

func AddSpaceMember(c *sdk.Client, spaceID, memberID, memberType, roleID string) (*sdk.WikiSpaceMember, error) {
	return c.AddWikiSpaceMember(spaceID, memberID, memberType, roleID)
}

func RemoveSpaceMember(c *sdk.Client, spaceID, memberID string) error {
	return c.RemoveWikiSpaceMember(spaceID, memberID)
}

func ListPages(c *sdk.Client, spaceID string) ([]sdk.WikiPage, error) {
	return c.ListWikiPages(spaceID, 0, 30)
}

func GetPage(c *sdk.Client, pageID string) (*sdk.WikiPage, error) {
	return c.GetWikiPage(pageID)
}

func CreatePage(c *sdk.Client, spaceID, parentID, title, pageType string) (*sdk.WikiPage, error) {
	return c.CreateWikiPage(spaceID, parentID, title, pageType)
}

func UpdatePage(c *sdk.Client, pageID, title string) (*sdk.WikiPage, error) {
	return c.UpdateWikiPage(pageID, title)
}

func DeletePage(c *sdk.Client, pageID string) error {
	return c.DeleteWikiPage(pageID)
}

func GetPageContent(c *sdk.Client, pageID string) (*sdk.WikiPageContent, error) {
	return c.GetWikiPageContent(pageID)
}

func UpdatePageContent(c *sdk.Client, pageID, content, format string) (*sdk.WikiPageContent, error) {
	return c.UpdateWikiPageContent(pageID, content, format)
}

func ListPageVersions(c *sdk.Client, pageID string) ([]sdk.WikiPageVersion, error) {
	return c.ListWikiPageVersions(pageID)
}
