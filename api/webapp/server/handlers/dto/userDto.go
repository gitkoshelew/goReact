package dto

import (
	"time"
)

// UserDto ...
type UserDto struct {
	UserID      int       `json:"userId"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string    `json:"role"`
	Verified    bool      `json:"verified"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	Sex         int       `json:"sex"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Photo       string    `json:"photo"`
}
