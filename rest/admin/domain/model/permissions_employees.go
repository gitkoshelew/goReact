package model

import validation "github.com/go-ozzo/ozzo-validation"

// PermissionsEmployees Model for linking two tables : permissions and employees
type PermissionsEmployees struct {
	Permissions Permission
	Employee    Employee
}

// PermissionsEmployeesDTO
type PermissionsEmployeesDTO struct {
	PermissionsID int
	EmployeeID    int
}

func (pe *PermissionsEmployeesDTO) Validate() error {
	return validation.ValidateStruct(
		pe,
		validation.Field(&pe.PermissionsID, validation.Required,validation.By(IsValidID)),
		validation.Field(&pe.EmployeeID, validation.Required, validation.By(IsValidID)),
	)
}
