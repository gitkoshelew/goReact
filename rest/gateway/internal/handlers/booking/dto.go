package booking

import "time"

// BookingDTO struct
type bookingDTO struct {
	BookingID  int       `json:"bookingId"`
	SeatID     int       `json:"seat"`
	PetID      int       `json:"pet"`
	EmployeeID int       `json:"employeeId"`
	Status     string    `json:"status"`
	StartDate  time.Time `json:"start"`
	EndDate    time.Time `json:"end"`
	Paid       bool      `json:"paid"`
	Notes      string    `json:"notes,omitempty"`
}

// DataValidation checks Seat, pet and Employee ids
type DataValidation struct {
	SeatID     int `json:"seat,omitempty"`
	PetID      int `json:"pet,omitempty"`
	EmployeeID int `json:"employeeId,omitempty"`
}
