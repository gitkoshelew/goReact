package store

import (
	"time"
)

// BookingStatus ...
type BookingStatus string

// BookingStatus options
var (
	BookingStatusPending    BookingStatus = "pending"
	BookingStatusInProgress BookingStatus = "in-progress"
	BookingStatusCompleted  BookingStatus = "completed"
	BookingStatusCancelled  BookingStatus = "cancelled"
)

// Booking struct
type Booking struct {
	BookingID   int           `json:"bookingId"`
	Seat        Seat          `json:"seat"`
	Pet         Pet           `json:"pet"`
	Employee    Employee      `json:"employeeId"`
	Status      BookingStatus `json:"status"`
	StartDate   time.Time     `json:"start"`
	EndDate     time.Time     `json:"end"`
	ClientNotes string        `json:"notes"`
}

// SetPet sets pet to booking
func (b *Booking) SetPet(p Pet) {
	b.Pet = p
}

// SetSeat sets Seat to booking
func (b *Booking) SetSeat(s Seat) {
	b.Seat = s
}

// SetStatus sets status to booking
func (b *Booking) SetStatus(bs BookingStatus) {
	b.Status = bs
}

// SetStartDate sets date of the start to booking
func (b *Booking) SetStartDate(t time.Time) {
	b.StartDate = t
}

// SetEndDate sets date of the end to booking
func (b *Booking) SetEndDate(t time.Time) {
	b.EndDate = t
}

// SetEmployee sets Employee to booking
func (b *Booking) SetEmployee(e Employee) {
	b.Employee = e
}

// SetClientNotes sets Client Notes to booking
func (b *Booking) SetClientNotes(s string) {
	b.ClientNotes = s
}
