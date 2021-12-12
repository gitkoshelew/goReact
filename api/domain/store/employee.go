package store

// Employee extends User and has all User (and Account) fields
type Employee struct {
	EmployeeID int `json:"employeeId"`
	User
	Hotel    Hotel
	Position string `json:"position"`
	Role     string `json:"role"`
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
