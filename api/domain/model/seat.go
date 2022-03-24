package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Seat struct
type Seat struct {
	SeatID      int `json:"seatId"`
	Room        Room
	Description string     `json:"description,omitempty"`
	RentFrom    *time.Time `json:"rentFrom,omitempty"`
	RentTo      *time.Time `json:"rentTo,omitempty"`
}

// SeatDTO struct
type SeatDTO struct {
	SeatID      int        `json:"seatId"`
	RoomID      int        `json:"roomId"`
	Description string     `json:"description,omitempty"`
	RentFrom    *time.Time `json:"rentFrom,omitempty"`
	RentTo      *time.Time `json:"rentTo,omitempty"`
}

// Validate ...
func (s *SeatDTO) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.RoomID, validation.Required, validation.By(IsValidID)),
		validation.Field(&s.Description, validation.By(IsSQL)),
		validation.Field(&s.RentFrom, validation.NotNil, validation.By(IsValidStartDate)),
		validation.Field(&s.RentTo, validation.NotNil, validation.By(IsValidEndDate)),
	)
}
