package entity

import "fmt"

// User ...
type User struct {
	Account
	UserID      int
	Name        string
	Surname     string
	MiddleName  string
	DateOfBirth string
	Address     string
	Phone       string
	Email       string
}

func (u *User) getInfo() string {
	return fmt.Sprintf("Account ID: %d\n"+
		"Login: %s\n"+
		"User ID: %d\n"+
		"Name: %s\n"+
		"Surname: %s\n"+
		"Middlename: %s\n"+
		"Date of birth: %s\n"+
		"Address: %s\n"+
		"Phone number: %s\n"+
		"Email: %s\n",
		u.AccountID, u.Login, u.UserID, u.Name, u.Surname, u.MiddleName, u.DateOfBirth, u.Address, u.Phone, u.Email)
}

func (u *User) setName(s string) {
	u.Name = s
}
func (u *User) setSurname(s string) {
	u.Surname = s
}
func (u *User) setMiddleName(s string) {
	u.MiddleName = s
}
func (u *User) setDateOfBirth(s string) {
	u.DateOfBirth = s
}
func (u *User) setAddress(s string) {
	u.Address = s
}
func (u *User) setPhone(s string) {
	u.Phone = s
}
func (u *User) setEmail(s string) {
	u.Email = s
}
