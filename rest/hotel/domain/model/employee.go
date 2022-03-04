package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Employee ...
type Employee struct {
	EmployeeID  int       `json:"employeeId,omitempty"`
	Email       string    `json:"email,omitempty"`
	Role        Role      `json:"role,omitempty"`
	Verified    bool      `json:"verified,omitempty"`
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"sName,omitempty"`
	MiddleName  string    `json:"mName,omitempty"`
	Sex         string    `json:"sex,omitempty"`
	DateOfBirth time.Time `json:"birthDate,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Photo       string    `json:"photo,omitempty"`
	Hotel       Hotel
	Position    Position `json:"position"`
}

// EmployeeDTO ...
type EmployeeDTO struct {
	EmployeeID  int       `json:"employeeId"`
	Email       string    `json:"email,omitempty"`
	Role        string    `json:"role,omitempty"`
	Verified    bool      `json:"verified,omitempty"`
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"sName,omitempty"`
	MiddleName  string    `json:"mName,omitempty"`
	Sex         string    `json:"sex,omitempty"`
	DateOfBirth time.Time `json:"birthDate,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Photo       string    `json:"photo,omitempty"`
	HotelID     int       `json:"hotelId,omitempty"`
	Position    string    `json:"position,omitempty"`
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

// Role ...
type Role string

// Role constants
const (
	ClientRole Role = "client"
)

// Validate ...
func (e *Employee) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.Hotel, validation.Required),
		validation.Field(&e.Position, validation.Required, validation.By(IsEmployeePosition)),
	)
}
