package entity

import "fmt"

// Employee ...
type Employee struct {
	User
	EmployeeID int
	Hotel      string
	Position   string
	Role       string
}

func (e *Employee) getInfo() string {
	return fmt.Sprintf("Account ID: %d\n"+
		"Login: %s\n"+
		"User ID: %d\n"+
		"Name: %s\n"+
		"Surname: %s\n"+
		"Middlename: %s\n"+
		"Date of birth: %s\n"+
		"Address: %s\n"+
		"Phone number: %s\n"+
		"Email: %s\n"+
		"Employee ID: %d\n"+
		"Hotel: %s\n"+
		"Position: %s\n"+
		"Role: %s\n",
		e.AccountID, e.Login, e.UserID, e.Name, e.Surname, e.MiddleName, e.DateOfBirth, e.Address, e.Phone, e.Email, e.EmployeeID, e.Hotel, e.Position, e.Role)
}

func (e *Employee) setHotel(s string) {
	e.Hotel = s
}

func (e *Employee) setPosition(s string) {
	e.Position = s
}

func (e *Employee) setRole(s string) {
	e.Role = s
}
