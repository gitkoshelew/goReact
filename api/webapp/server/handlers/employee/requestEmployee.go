package employee

type employeeRequest struct {
	EmployeeID int    `json:"employeeId"`
	UserID     int    `json:"userId"`
	HotelID    int    `json:"hotelId"`
	Position   string `json:"position"`
	Role       string `json:"role"`
}
