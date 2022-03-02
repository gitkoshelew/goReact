package model

import "time"

// Booking struct
type Booking struct {
	BookingID  int           `json:"bookingId"`
	SeatID     int           `json:"seat"`
	PetID      int           `json:"pet"`
	EmployeeID int           `json:"employeeId"`
	Status     BookingStatus `json:"status"`
	StartDate  *time.Time    `json:"start"`
	EndDate    *time.Time    `json:"end"`
	Paid       *bool         `json:"paid"`
	Notes      string        `json:"notes,omitempty"`
}

// BookingStatus ...
type BookingStatus string

// BookingStatus options
var (
	BookingStatusPending    BookingStatus = "pending"
	BookingStatusInProgress BookingStatus = "in-progress"
	BookingStatusCompleted  BookingStatus = "completed"
	BookingStatusCancelled  BookingStatus = "cancelled"
)
