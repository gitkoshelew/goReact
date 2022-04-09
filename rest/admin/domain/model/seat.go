package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Seat struct
type Seat struct {
	SeatID   int `json:"seatId"`
	Room     Room
	RentFrom *time.Time `json:"rentFrom,omitempty"`
	RentTo   *time.Time `json:"rentTo,omitempty"`
	Price    float64    `json:"price"`
}

// SeatDTO struct
type SeatDTO struct {
	SeatID   int        `json:"seatId"`
	RoomID   int        `json:"roomId"`
	RentFrom *time.Time `json:"rentFrom,omitempty"`
	RentTo   *time.Time `json:"rentTo,omitempty"`
	Price    float64    `json:"price"`
}

// Validate ...
func (s *SeatDTO) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.RoomID, validation.Required, validation.By(IsValidID)),
		validation.Field(&s.RentFrom, validation.NotNil, validation.By(IsValidStartDate)),
		validation.Field(&s.RentTo, validation.NotNil, validation.By(IsValidEndDate)),
		validation.Field(&s.Price, validation.Required, validation.Min(0.01), validation.Max(9999999999.9)),
	)
}
