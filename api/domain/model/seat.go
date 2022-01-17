package model

// Seat struct
type Seat struct {
	SeatID      int    `json:"seatId"`
	Description string `json:"description"`
	IsFree      bool   `json:"isFree"`
	Room        Room
}
