package entity

// Employee extends User and has all User (and Account) fields
type Employee struct {
	User
	EmployeeID int
	Hotel      Hotel
	Position   string
	Role       string
}

// SetHotel sets Employyes Hotel
func (e *Employee) SetHotel(h Hotel) {
	e.Hotel = h
}

// SetPosition sets Employyes Position
func (e *Employee) SetPosition(s string) {
	e.Position = s
}

// SetRole sets Employyes Role
func (e *Employee) SetRole(s string) {
	e.Role = s
}
