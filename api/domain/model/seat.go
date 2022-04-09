package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Seat struct
type Seat struct {
	SeatID      int `json:"seatId"`
	Room        Room
	Description string `json:"description,omitempty"`
	Price    float64    `json:"price"`
}

// SeatDTO struct
type SeatDTO struct {
	SeatID      int    `json:"seatId"`
	RoomID      int    `json:"roomId"`
	Description string `json:"description,omitempty"`
	Price    float64    `json:"price"`
}

// Validate ...
func (s *SeatDTO) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.RoomID, validation.Required, validation.By(IsValidID)),
		validation.Field(&s.Description, validation.By(IsSQL)),
		validation.Field(&s.Price, validation.Required, validation.Min(0.01), validation.Max(9999999999.9)),
	)
}
