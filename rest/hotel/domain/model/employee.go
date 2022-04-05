package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Employee ...
type Employee struct {
	EmployeeID  int        `json:"employeeId,omitempty"`
	Email       string     `json:"email,omitempty"`
	Role        Role       `json:"role,omitempty"`
	Verified    *bool      `json:"verified,omitempty"`
	Name        string     `json:"name,omitempty"`
	Surname     string     `json:"sName,omitempty"`
	MiddleName  string     `json:"mName,omitempty"`
	Sex         Sex        `json:"sex,omitempty"`
	DateOfBirth *time.Time `json:"birthDate,omitempty"`
	Address     string     `json:"address,omitempty"`
	Phone       string     `json:"phone,omitempty"`
	Photo       string     `json:"photo,omitempty"`
	Hotel       Hotel
	Position    Position `json:"position"`
}

// EmployeeDTO ...
type EmployeeDTO struct {
	EmployeeID  int        `json:"employeeId"`
	Email       string     `json:"email,omitempty"`
	Role        string     `json:"role,omitempty"`
	Verified    *bool      `json:"verified,omitempty"`
	Name        string     `json:"name,omitempty"`
	Surname     string     `json:"sName,omitempty"`
	MiddleName  string     `json:"mName,omitempty"`
	Sex         string     `json:"sex,omitempty"`
	DateOfBirth *time.Time `json:"birthDate,omitempty"`
	Address     string     `json:"address,omitempty"`
	Phone       string     `json:"phone,omitempty"`
	Photo       string     `json:"photo,omitempty"`
	HotelID     int        `json:"hotelId,omitempty"`
	Position    string     `json:"position,omitempty"`
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
	ClientRole   Role = "client"
	EmployeeRole Role = "employee"
)

// Sex ...
type Sex string

// Sex constants
const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

// Validate ...
func (e *EmployeeDTO) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.Email, validation.Required, is.Email, validation.By(IsSQL)),
		validation.Field(&e.Role, validation.By(IsRole)),
		validation.Field(&e.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&e.Surname, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&e.MiddleName, validation.By(IsLetterHyphenSpaces), validation.Length(0, 30), validation.By(IsSQL)),
		validation.Field(&e.Sex, validation.Required, validation.By(IsSex)),
		validation.Field(&e.DateOfBirth, validation.Required, validation.By(IsValidBirthDate)),
		validation.Field(&e.Address, validation.Required, validation.By(IsSQL), validation.Length(10, 40)),
		validation.Field(&e.Phone, validation.Required, validation.By(IsPhone)),
		validation.Field(&e.Photo, validation.By(IsSQL)),
		validation.Field(&e.HotelID, validation.Required, validation.By(IsValidID)),
		validation.Field(&e.Position, validation.Required, validation.By(IsEmployeePosition)),
	)
}
