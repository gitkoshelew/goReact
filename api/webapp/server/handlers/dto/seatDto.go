package dto

// SeatDto ...
type SeatDto struct {
	HotelRoomSeatID int    `json:"seatId"`
	Description     string `json:"desc"`
	IsFree          bool   `json:"isFree"`
}
