package model

// PermissionsEmployees Model for linking two tables : permissions and employees
type PermissionsEmployees struct {
	Permissions Permission
	Employee    Employee
}
