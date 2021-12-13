package dto

import (
	"time"
)

// UserDto ...
type UserDto struct {
	AccountID   int       `json:"accountId"`
	UserID      int       `json:"userId"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
}
