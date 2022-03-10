package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Booking struct
type Booking struct {
	BookingID int `json:"bookingId"`
	Seat
	Pet
	Employee
	Status        BookingStatus `json:"status"`
	StartDate     *time.Time    `json:"start"`
	EndDate       *time.Time    `json:"end"`
	Notes         string        `json:"notes,omitempty"`
	TransactionID int           `json:"transactionId"`
	Paid          *bool         `json:"paid"`
}

// BookingDTO struct
type BookingDTO struct {
	BookingID     int           `json:"bookingId"`
	SeatID        int           `json:"seat"`
	PetID         int           `json:"pet"`
	EmployeeID    int           `json:"employeeId"`
	Status        BookingStatus `json:"status"`
	StartDate     *time.Time    `json:"start"`
	EndDate       *time.Time    `json:"end"`
	Notes         string        `json:"notes,omitempty"`
	TransactionID int           `json:"transactionId"`
	Paid          *bool         `json:"paid"`
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

// Validate ...
func (b *Booking) Validate() error {
	return validation.ValidateStruct(
		b,
		validation.Field(&b.Seat, validation.Required),
		validation.Field(&b.Pet, validation.Required),
		validation.Field(&b.Employee, validation.Required),
		validation.Field(&b.Status, validation.Required, validation.By(IsBookingStatus)),
		validation.Field(&b.StartDate, validation.Required),
		validation.Field(&b.EndDate, validation.Required),
		validation.Field(&b.Paid, validation.NotNil),
	)
}
