package testhub

import (
	"pingcode-client/internal/pkg/sdk"
)

func ListLibraries(client *sdk.Client) ([]sdk.TestLibrary, error) {
	return client.ListTestLibraries()
}

func CreateLibrary(client *sdk.Client, name, identifier, scopeType, scopeID, visibility, desc string) (*sdk.TestLibrary, error) {
	return client.CreateTestLibrary(name, identifier, scopeType, scopeID, visibility, desc)
}

func UpdateLibrary(client *sdk.Client, libraryID, name, visibility, desc string) (*sdk.TestLibrary, error) {
	return client.UpdateTestLibrary(libraryID, name, visibility, desc)
}

func ListLibraryMembers(client *sdk.Client, libraryID string) ([]sdk.TestLibraryMember, error) {
	return client.ListTestLibraryMembers(libraryID)
}

func AddLibraryMember(client *sdk.Client, libraryID, memberID, memberType, roleID string) (*sdk.TestLibraryMember, error) {
	return client.AddTestLibraryMember(libraryID, memberID, memberType, roleID)
}

func RemoveLibraryMember(client *sdk.Client, libraryID, memberID string) error {
	return client.RemoveTestLibraryMember(libraryID, memberID)
}

func ListSuites(client *sdk.Client, libraryID string) ([]sdk.TestSuite, error) {
	return client.ListTestSuites(libraryID)
}

func CreateSuite(client *sdk.Client, libraryID, name string) (*sdk.TestSuite, error) {
	return client.CreateTestSuite(libraryID, name)
}

func DeleteSuite(client *sdk.Client, libraryID, suiteID string) error {
	return client.DeleteTestSuite(libraryID, suiteID)
}

func ListCases(client *sdk.Client, libraryID, suiteID string) ([]sdk.TestCase, error) {
	return client.ListTestCases(libraryID, suiteID)
}

func GetCase(client *sdk.Client, caseID string) (*sdk.TestCase, error) {
	return client.GetTestCase(caseID)
}

func CreateCase(client *sdk.Client, libraryID, title, suiteID, typeID, maintenanceID, desc, precondition string) (*sdk.TestCase, error) {
	tc := &sdk.TestCase{
		LibraryID:     libraryID,
		Title:         title,
		SuiteID:       suiteID,
		TypeID:        typeID,
		MaintenanceID: maintenanceID,
		Description:   desc,
		Precondition:  precondition,
	}
	return client.CreateTestCase(tc)
}

func UpdateCase(client *sdk.Client, caseID, title, suiteID, typeID, maintenanceID, desc, precondition string) (*sdk.TestCase, error) {
	tc := &sdk.TestCase{
		Title:         title,
		SuiteID:       suiteID,
		TypeID:        typeID,
		MaintenanceID: maintenanceID,
		Description:   desc,
		Precondition:  precondition,
	}
	return client.UpdateTestCase(caseID, tc)
}

func DeleteCase(client *sdk.Client, caseID string) error {
	return client.DeleteTestCase(caseID)
}

func ListPlans(client *sdk.Client, libraryID string) ([]sdk.TestPlan, error) {
	return client.ListTestPlans(libraryID)
}

func CreatePlan(client *sdk.Client, libraryID, name, typeID, assigneeID string, startAt, endAt int64) (*sdk.TestPlan, error) {
	plan := &sdk.TestPlan{
		Name:       name,
		TypeID:     typeID,
		AssigneeID: assigneeID,
		StartAt:    startAt,
		EndAt:      endAt,
	}
	return client.CreateTestPlan(libraryID, plan)
}

func CreateRun(client *sdk.Client, libraryID, planID, caseID, executorID string) (*sdk.TestRun, error) {
	return client.CreateTestRun(libraryID, planID, caseID, executorID)
}
