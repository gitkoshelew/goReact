package model

import validation "github.com/go-ozzo/ozzo-validation"

// Employee extends User and has all User fields
type Employee struct {
	EmployeeID int `json:"employeeId"`
	User
	Hotel    Hotel
	Position Position `json:"position"`
}

// Position ...
type Position string

// Position constants
const (
	ManagerPosition  Position = "manager"
	EmployeePosition Position = "employee"
	OwnerPosition    Position = "owner"
	AdminPosition    Position = "admin"
)


func (e Employee)  PositionString() string{
	return string(e.Position)
}

// Validate ...
func (e *Employee) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.User, validation.Required),
		validation.Field(&e.Hotel, validation.Required),
		validation.Field(&e.Position, validation.Required, validation.By(IsEmployeePosition)),
	)
}
