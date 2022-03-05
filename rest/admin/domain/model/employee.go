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
var (
	ManagerPosition  Position = "manager"
	EmployeePosition Position = "employee"
	OwnerPosition    Position = "owner"
	AdminPosition    Position = "admin"
)

func (e Employee) PositionString() string {
	switch e.Position {
	case ManagerPosition:
		return "manager"
	case EmployeePosition:
		return "employee"
	case OwnerPosition:
		return "owner"
	case AdminPosition:
		return "admin"
	}
	return "unknown"
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
