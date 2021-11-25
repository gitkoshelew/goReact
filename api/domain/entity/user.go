package entity

import (
	"goReact/pkg/date"
)

// User extends Account and has all Account fields
type User struct {
	Account
	UserID      int
	Name        string
	Surname     string
	MiddleName  string
	DateOfBirth date.Date
	Address     string
	Phone       string
	Email       string
}

// SetName sets Users name
func (u *User) SetName(s string) {
	u.Name = s
}

// SetSurname sets Users Surname
func (u *User) SetSurname(s string) {
	u.Surname = s
}

// SetMiddleName sets Users Middlename
func (u *User) SetMiddleName(s string) {
	u.MiddleName = s
}

// SetDateOfBirth sets Users date of birth
func (u *User) SetDateOfBirth(d date.Date) {
	u.DateOfBirth = d
}

// SetAddress sets Users Address
func (u *User) SetAddress(s string) {
	u.Address = s
}

// SetPhone sets Users phone number
func (u *User) SetPhone(s string) {
	u.Phone = s
}

// SetEmail sets Users email
func (u *User) SetEmail(s string) {
	u.Email = s
}
