package dto

// Employee DTO
type Employee struct {
	UserID     int    `json:"userId"`
	HotelID    int    `json:"hotelId"`
	EmployeeID int    `json:"employeeId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
}
