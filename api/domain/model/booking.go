package model

import (
	"time"
)

// Booking struct
type Booking struct {
	BookingID int           `json:"bookingId"`
	Seat      Seat          `json:"seat"`
	Pet       Pet           `json:"pet"`
	Employee  Employee      `json:"employeeId"`
	Status    BookingStatus `json:"status"`
	StartDate time.Time     `json:"start"`
	EndDate   time.Time     `json:"end"`
	Notes     string        `json:"notes"`
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
