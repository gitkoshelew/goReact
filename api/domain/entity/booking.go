package entity

import (
	"goReact/domain/dto"
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
	BookingID   int           `json:"bookingId"`
	Pet         Pet           `json:"pet"`
	Seat        HotelRoomSeat `json:"seat"`
	Status      BookingStatus `json:"status"`
	StartDate   date.Date     `json:"start"`
	EndDate     date.Date     `json:"end"`
	Employee    Employee      `json:"employeeId"`
	ClientNotes string        `json:"notes"`
}

// SetPet sets pet to booking
func (b *Booking) SetPet(p Pet) {
	b.Pet = p
}

// SetSeat sets Seat to booking
func (b *Booking) SetSeat(h HotelRoomSeat) {
	b.Seat = h
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

// GetBookingByID returns Booking by Id from storage
func GetBookingByID(id int) Booking {
	var booking Booking
	for _, bk := range GetBookings() {
		if id == bk.BookingID {
			booking = bk
			break
		}
	}
	return booking
}

// GetBookingsByID returns []Bookings by []ids from storage
func GetBookingsByID(ids []int) []Booking {
	var bookings []Booking
	for _, id := range ids {
		bookings = append(bookings, GetBookingByID(id))
	}
	return bookings
}

// BookingToDto makes DTO from Booking object
func BookingToDto(b Booking) dto.Booking {
	return dto.Booking{
		BookingID:   b.BookingID,
		PetID:       b.Pet.PetID,
		SeatID:      b.Seat.HotelRoomSeatID,
		Status:      string(b.Status),
		StartDate:   b.StartDate,
		EndDate:     b.EndDate,
		EmployeeID:  b.Employee.EmployeeID,
		ClientNotes: b.ClientNotes,
	}
}
