package user

import (
	"time"
)

// Request ...
type Request struct {
	UserID      int       `json:"userId"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Verified    bool      `json:"verified"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	Sex         string    `json:"sex"`
	DateOfBirth time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Photo       string    `json:"photo"`
}
