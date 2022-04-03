package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
	Verified *bool  `json:"verified"`
}

// UserDTO ...
type UserDTO struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Verified *bool  `json:"verified"`
}

// Role ...
type Role string

// Role constants
const (
	ClientRole    Role = "client"
	EmployeeRole  Role = "employee"
	AnonymousRole Role = "anonymous"
)

// Sex ...
type Sex string

// Sex constants
const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

// Validate ...
func (u *UserDTO) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.By(IsSQL), validation.Length(5, 100)),
		validation.Field(&u.Role, validation.Required, validation.By(IsRole)),
	)
}
