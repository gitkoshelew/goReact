package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Employee extends User and has all User fields
type Employee struct {
	EmployeeID int `json:"employeeId"`
	User
	Hotel    Hotel
	Position Position `json:"position"`
}

// EmployeeDTO ...
type EmployeeDTO struct {
	EmployeeID int    `json:"employeeId"`
	UserID     int    `json:"userId"`
	HotelID    int    `json:"hotelId"`
	Position   string `json:"position"`
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

// Validate ...
func (e *EmployeeDTO) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.UserID, validation.Required, validation.By(IsValidID)),
		validation.Field(&e.HotelID, validation.Required, validation.By(IsValidID)),
		validation.Field(&e.Position, validation.Required, validation.By(IsEmployeePosition)),
	)
}
