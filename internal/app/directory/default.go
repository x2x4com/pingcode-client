package directory

import (
	"fmt"
	"log"
	"pingcode-client/internal/pkg/sdk"
)

func ListUsers(client *sdk.Client, keywords, name, departmentIDs string) {
	users, err := client.ListDirectoryUsers(keywords, name, departmentIDs)
	if err != nil {
		log.Fatalf("Error listing directory users: %v", err)
	}
	fmt.Println("Directory Users:")
	for _, u := range users {
		fmt.Printf("- %s / %s (ID: %s)\n", u.Name, u.DisplayName, u.ID)
	}
}

func CreateUser(client *sdk.Client, name, displayName, email, mobile, password, departmentID, jobID, employeeNumber string) {
	if name == "" || displayName == "" {
		log.Fatal("Name and display name are required")
	}
	u := &sdk.DirectoryUser{
		Name:           name,
		DisplayName:    displayName,
		Email:          email,
		Mobile:         mobile,
		Password:       password,
		DepartmentID:   departmentID,
		JobID:          jobID,
		EmployeeNumber: employeeNumber,
	}
	created, err := client.CreateDirectoryUser(u)
	if err != nil {
		log.Fatalf("Error creating directory user: %v", err)
	}
	fmt.Printf("Created User: %s / %s (ID: %s)\n", created.Name, created.DisplayName, created.ID)
}

func UpdateUser(client *sdk.Client, userID, name, displayName, email, mobile, status, departmentID, jobID, employeeNumber string) {
	if userID == "" {
		log.Fatal("User ID is required")
	}
	u := &sdk.DirectoryUser{
		Name:           name,
		DisplayName:    displayName,
		Email:          email,
		Mobile:         mobile,
		Status:         status,
		DepartmentID:   departmentID,
		JobID:          jobID,
		EmployeeNumber: employeeNumber,
	}
	updated, err := client.UpdateDirectoryUser(userID, u)
	if err != nil {
		log.Fatalf("Error updating directory user: %v", err)
	}
	fmt.Printf("Updated User: %s / %s (ID: %s)\n", updated.Name, updated.DisplayName, updated.ID)
}

func ListDepartments(client *sdk.Client) {
	depts, err := client.ListDepartments()
	if err != nil {
		log.Fatalf("Error listing departments: %v", err)
	}
	fmt.Println("Departments:")
	for _, d := range depts {
		fmt.Printf("- %s (ID: %s)\n", d.Name, d.ID)
	}
}

func CreateDepartment(client *sdk.Client, name, parentID, headID string) {
	if name == "" {
		log.Fatal("Department name is required")
	}
	d, err := client.CreateDepartment(name, parentID, headID)
	if err != nil {
		log.Fatalf("Error creating department: %v", err)
	}
	fmt.Printf("Created Department: %s (ID: %s)\n", d.Name, d.ID)
}

func UpdateDepartment(client *sdk.Client, departmentID, name, parentID, headID string) {
	if departmentID == "" {
		log.Fatal("Department ID is required")
	}
	d, err := client.UpdateDepartment(departmentID, name, parentID, headID)
	if err != nil {
		log.Fatalf("Error updating department: %v", err)
	}
	fmt.Printf("Updated Department: %s (ID: %s)\n", d.Name, d.ID)
}

func DeleteDepartment(client *sdk.Client, departmentID string) {
	if departmentID == "" {
		log.Fatal("Department ID is required")
	}
	err := client.DeleteDepartment(departmentID)
	if err != nil {
		log.Fatalf("Error deleting department: %v", err)
	}
	fmt.Printf("Deleted department %s\n", departmentID)
}
