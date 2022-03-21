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
	RentFrom    *time.Time `json:"rentFrom"`
	RentTo      *time.Time `json:"rentTo"`
}

// SeatDTO struct
type SeatDTO struct {
	SeatID      int        `json:"seatId"`
	RoomID      int        `json:"roomId"`
	Description string     `json:"description,omitempty"`
	RentFrom    *time.Time `json:"rentFrom"`
	RentTo      *time.Time `json:"rentTo"`
}

// Validate ...
func (s *SeatDTO) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.RoomID, validation.Required, validation.By(IsValidID)),
		validation.Field(&s.Description, validation.Required, validation.By(IsSQL)),
		validation.Field(&s.RentFrom, validation.Required, validation.By(IsValidStartDate)),
		validation.Field(&s.RentTo, validation.Required, validation.By(IsValidEndDate)),
	)
}
