package auth

import "time"

// UserDTO ...
type UserDTO struct {
	UserID      int       `json:"userId,omitempty"`
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
}
