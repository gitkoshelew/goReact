package entity

import (
	"goReact/pkg/date"
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
	BookingID   int
	Pet         *Pet
	Seat        *HotelRoomSeat
	Status      BookingStatus
	StartDate   date.Date
	EndDate     date.Date
	Employee    Employee
	ClientNotes string
}

// SetPet sets pet to booking
func (b *Booking) SetPet(p Pet) {
	b.Pet = &p
}

// SetSeat sets Seat to booking
func (b *Booking) SetSeat(h HotelRoomSeat) {
	b.Seat = &h
}

// SetStatus sets status to booking
func (b *Booking) SetStatus(bs BookingStatus) {
	b.Status = bs
}

// SetStartDate sets date of the start to booking
func (b *Booking) SetStartDate(d date.Date) {
	b.StartDate = d
}

// SetEndDate sets date of the end to booking
func (b *Booking) SetEndDate(d date.Date) {
	b.EndDate = d
}

// SetEmployee sets Employee to booking
func (b *Booking) SetEmployee(e Employee) {
	b.Employee = e
}

// SetClientNotes sets Client Notes to booking
func (b *Booking) SetClientNotes(s string) {
	b.ClientNotes = s
}
