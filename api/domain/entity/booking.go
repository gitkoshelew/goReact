package entity

import (
	"goReact/pkg/date"
)

type BookingId string

type BookingStatus string

var (
	BookingStatusPending    BookingStatus = "pending"
	BookingStatusInProgress BookingStatus = "in-progress"
	BookingStatusCompleted  BookingStatus = "completed"
	BookingStatusCancelled  BookingStatus = "cancelled"
)

type Booking struct {
	Id        BookingId
	Pet       *Pet
	Seat      *HotelRoomSeat
	Status    BookingStatus
	StartDate date.Date
	EndDate   date.Date
	UserNotes string
}
