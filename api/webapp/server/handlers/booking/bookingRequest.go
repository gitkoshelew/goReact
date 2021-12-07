package booking

import "goReact/pkg/date"

type bookingRequest struct {
	BookingID   int       `json:"bookingId"`
	PetID       int       `json:"petId"`
	SeatID      int       `json:"seatId"`
	Status      string    `json:"status"`
	StartDate   date.Date `json:"start"`
	EndDate     date.Date `json:"end"`
	EmployeeID  int       `json:"employeeId"`
	ClientNotes string    `json:"notes"`
}
