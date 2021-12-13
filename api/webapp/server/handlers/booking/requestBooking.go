package booking

import (
	"time"
)

type bookingRequest struct {
	BookingID   int       `json:"bookingId"`
	PetID       int       `json:"petId"`
	SeatID      int       `json:"seatId"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start"`
	EndDate     time.Time `json:"end"`
	EmployeeID  int       `json:"employeeId"`
	ClientNotes string    `json:"notes"`
}
