package model

import validation "github.com/go-ozzo/ozzo-validation"

// Employee ...
type Employee struct {
	EmployeeID int `json:"employeeId"`
	UserID     int `json:"userId"`
	Hotel      Hotel
	Position   Position `json:"position"`
}

// EmployeeDTO ...
type EmployeeDTO struct {
	EmployeeID int      `json:"employeeId"`
	UserID     int      `json:"userId"`
	HotelID    int      `json:"hotelId"`
	Position   Position `json:"position"`
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

// PositionString ...
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
		validation.Field(&e.UserID, validation.Required, validation.Min(1)),
		validation.Field(&e.Hotel, validation.Required),
		validation.Field(&e.Position, validation.Required, validation.By(IsEmployeePosition)),
	)
}
