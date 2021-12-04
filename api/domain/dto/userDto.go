package dto

import "goReact/pkg/date"

// User DTO
type User struct {
	AccountID   int       `json:"accountId"`
	UserID      int       `json:"userId"`
	Name        string    `json:"name"`
	Surname     string    `json:"sName"`
	MiddleName  string    `json:"mName"`
	DateOfBirth date.Date `json:"birthDate"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
}
