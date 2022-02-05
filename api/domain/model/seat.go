package model

import "time"

// Seat struct
type Seat struct {
	SeatID      int `json:"seatId"`
	Room        Room
	Description string    `json:"description,omitempty"`
	RentFrom    time.Time `json:"rentFrom"`
	RentTo      time.Time `json:"rentTo"`
}

// SeatDTO struct
type SeatDTO struct {
	SeatID      int       `json:"seatId"`
	RoomID      int       `json:"roomId"`
	Description string    `json:"description,omitempty"`
	RentFrom    time.Time `json:"rentFrom"`
	RentTo      time.Time `json:"rentTo"`
}
