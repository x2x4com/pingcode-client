package testhub

import (
	"fmt"
	"log"
	"pingcode-client/internal/pkg/sdk"
)

func ListLibraries(client *sdk.Client) {
	libs, err := client.ListTestLibraries()
	if err != nil {
		log.Fatalf("Error listing test libraries: %v", err)
	}
	fmt.Println("Test Libraries:")
	for _, l := range libs {
		fmt.Printf("- %s (ID: %s)\n", l.Name, l.ID)
	}
}

func CreateLibrary(client *sdk.Client, name, identifier, scopeType, scopeID, visibility, desc string) {
	lib, err := client.CreateTestLibrary(name, identifier, scopeType, scopeID, visibility, desc)
	if err != nil {
		log.Fatalf("Error creating test library: %v", err)
	}
	fmt.Printf("Created Test Library: %s (ID: %s)\n", lib.Name, lib.ID)
}

func UpdateLibrary(client *sdk.Client, libraryID, name, visibility, desc string) {
	lib, err := client.UpdateTestLibrary(libraryID, name, visibility, desc)
	if err != nil {
		log.Fatalf("Error updating test library: %v", err)
	}
	fmt.Printf("Updated Test Library: %s (ID: %s)\n", lib.Name, lib.ID)
}

func ListLibraryMembers(client *sdk.Client, libraryID string) {
	members, err := client.ListTestLibraryMembers(libraryID)
	if err != nil {
		log.Fatalf("Error listing test library members: %v", err)
	}
	fmt.Printf("Test Library Members (library: %s):\n", libraryID)
	for _, m := range members {
		fmt.Printf("- %s (type: %s)\n", m.ID, m.Type)
	}
}

func AddLibraryMember(client *sdk.Client, libraryID, memberID, memberType, roleID string) {
	m, err := client.AddTestLibraryMember(libraryID, memberID, memberType, roleID)
	if err != nil {
		log.Fatalf("Error adding test library member: %v", err)
	}
	fmt.Printf("Added Member: %s (type: %s) to Library: %s\n", m.ID, m.Type, libraryID)
}

func RemoveLibraryMember(client *sdk.Client, libraryID, memberID string) {
	if err := client.RemoveTestLibraryMember(libraryID, memberID); err != nil {
		log.Fatalf("Error removing test library member: %v", err)
	}
	fmt.Printf("Removed Member: %s from Library: %s\n", memberID, libraryID)
}

func ListSuites(client *sdk.Client, libraryID string) {
	suites, err := client.ListTestSuites(libraryID)
	if err != nil {
		log.Fatalf("Error listing test suites: %v", err)
	}
	fmt.Printf("Test Suites (library: %s):\n", libraryID)
	for _, s := range suites {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func CreateSuite(client *sdk.Client, libraryID, name string) {
	suite, err := client.CreateTestSuite(libraryID, name)
	if err != nil {
		log.Fatalf("Error creating test suite: %v", err)
	}
	fmt.Printf("Created Test Suite: %s (ID: %s)\n", suite.Name, suite.ID)
}

func DeleteSuite(client *sdk.Client, libraryID, suiteID string) {
	if err := client.DeleteTestSuite(libraryID, suiteID); err != nil {
		log.Fatalf("Error deleting test suite: %v", err)
	}
	fmt.Printf("Deleted Test Suite: %s from Library: %s\n", suiteID, libraryID)
}

func ListCases(client *sdk.Client, libraryID string) {
	cases, err := client.ListTestCases(libraryID)
	if err != nil {
		log.Fatalf("Error listing test cases: %v", err)
	}
	fmt.Printf("Test Cases (library: %s):\n", libraryID)
	for _, tc := range cases {
		fmt.Printf("- %s (ID: %s)\n", tc.Title, tc.ID)
	}
}

func GetCase(client *sdk.Client, caseID string) {
	tc, err := client.GetTestCase(caseID)
	if err != nil {
		log.Fatalf("Error getting test case: %v", err)
	}
	fmt.Printf("ID:            %s\nTitle:         %s\nLibraryID:     %s\nSuiteID:       %s\nTypeID:        %s\nMaintenanceID: %s\nDescription:   %s\nPrecondition:  %s\n",
		tc.ID, tc.Title, tc.LibraryID, tc.SuiteID, tc.TypeID, tc.MaintenanceID, tc.Description, tc.Precondition)
}

func CreateCase(client *sdk.Client, libraryID, title, suiteID, typeID, maintenanceID, desc, precondition string) {
	tc := &sdk.TestCase{
		LibraryID:     libraryID,
		Title:         title,
		SuiteID:       suiteID,
		TypeID:        typeID,
		MaintenanceID: maintenanceID,
		Description:   desc,
		Precondition:  precondition,
	}
	created, err := client.CreateTestCase(tc)
	if err != nil {
		log.Fatalf("Error creating test case: %v", err)
	}
	fmt.Printf("Created Test Case: %s (ID: %s)\n", created.Title, created.ID)
}

func UpdateCase(client *sdk.Client, caseID, title, suiteID, typeID, maintenanceID, desc, precondition string) {
	tc := &sdk.TestCase{
		Title:         title,
		SuiteID:       suiteID,
		TypeID:        typeID,
		MaintenanceID: maintenanceID,
		Description:   desc,
		Precondition:  precondition,
	}
	updated, err := client.UpdateTestCase(caseID, tc)
	if err != nil {
		log.Fatalf("Error updating test case: %v", err)
	}
	fmt.Printf("Updated Test Case: %s (ID: %s)\n", updated.Title, updated.ID)
}

func DeleteCase(client *sdk.Client, caseID string) {
	if err := client.DeleteTestCase(caseID); err != nil {
		log.Fatalf("Error deleting test case: %v", err)
	}
	fmt.Printf("Deleted Test Case: %s\n", caseID)
}

func ListPlans(client *sdk.Client, libraryID string) {
	plans, err := client.ListTestPlans(libraryID)
	if err != nil {
		log.Fatalf("Error listing test plans: %v", err)
	}
	fmt.Printf("Test Plans (library: %s):\n", libraryID)
	for _, p := range plans {
		fmt.Printf("- %s (ID: %s)\n", p.Name, p.ID)
	}
}

func CreatePlan(client *sdk.Client, libraryID, name, typeID, assigneeID string, startAt, endAt int64) {
	plan := &sdk.TestPlan{
		Name:       name,
		TypeID:     typeID,
		AssigneeID: assigneeID,
		StartAt:    startAt,
		EndAt:      endAt,
	}
	created, err := client.CreateTestPlan(libraryID, plan)
	if err != nil {
		log.Fatalf("Error creating test plan: %v", err)
	}
	fmt.Printf("Created Test Plan: %s (ID: %s)\n", created.Name, created.ID)
}

func CreateRun(client *sdk.Client, libraryID, planID, caseID, executorID string) {
	run, err := client.CreateTestRun(libraryID, planID, caseID, executorID)
	if err != nil {
		log.Fatalf("Error creating test run: %v", err)
	}
	fmt.Printf("Created Test Run: %s (library: %s, plan: %s, case: %s)\n", run.ID, run.LibraryID, run.PlanID, run.CaseID)
}
