package entity

// Employee extends User and has all User (and Account) fields
type Employee struct {
	User
	Hotel      Hotel
	EmployeeID int    `json:"employeeId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
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

// GetEmployeeByID returns Employee by id from storage
func GetEmployeeByID(id int) Employee {
	employees := GetEmployees()
	var employee Employee
	for _, e := range employees {
		if id == e.EmployeeID {
			employee = e
		}
	}
	return employee
}
