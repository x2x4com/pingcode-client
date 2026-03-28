package directory

import (
	"pingcode-client/internal/pkg/sdk"
)

func ListUsers(client *sdk.Client, keywords, name, departmentIDs string) ([]sdk.DirectoryUser, error) {
	return client.ListDirectoryUsers(keywords, name, departmentIDs)
}

func CreateUser(client *sdk.Client, name, displayName, email, mobile, password, departmentID, jobID, employeeNumber string) (*sdk.DirectoryUser, error) {
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
	return client.CreateDirectoryUser(u)
}

func UpdateUser(client *sdk.Client, userID, name, displayName, email, mobile, status, departmentID, jobID, employeeNumber string) (*sdk.DirectoryUser, error) {
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
	return client.UpdateDirectoryUser(userID, u)
}

func ListDepartments(client *sdk.Client) ([]sdk.Department, error) {
	return client.ListDepartments()
}

func CreateDepartment(client *sdk.Client, name, parentID, headID string) (*sdk.Department, error) {
	return client.CreateDepartment(name, parentID, headID)
}

func UpdateDepartment(client *sdk.Client, departmentID, name, parentID, headID string) (*sdk.Department, error) {
	return client.UpdateDepartment(departmentID, name, parentID, headID)
}

func DeleteDepartment(client *sdk.Client, departmentID string) error {
	return client.DeleteDepartment(departmentID)
}
