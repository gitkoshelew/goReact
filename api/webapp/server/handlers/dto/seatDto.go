package dto

// SeatDto ...
type SeatDto struct {
	SeatID      int    `json:"seatId"`
	RoomID      int    `json:"roomId"`
	IsFree      bool   `json:"isFree"`
	Description string `json:"desc"`
}
