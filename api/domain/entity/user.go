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

// set Users name
func (u *User) setName(s string) {
	u.Name = s
}

// set Users Surname
func (u *User) setSurname(s string) {
	u.Surname = s
}

// set Users Middlename
func (u *User) setMiddleName(s string) {
	u.MiddleName = s
}

// set Users date of birth
func (u *User) setDateOfBirth(d date.Date) {
	u.DateOfBirth = d
}

// set Users Address
func (u *User) setAddress(s string) {
	u.Address = s
}

// set Users phone number
func (u *User) setPhone(s string) {
	u.Phone = s
}

// set Users email
func (u *User) setEmail(s string) {
	u.Email = s
}
