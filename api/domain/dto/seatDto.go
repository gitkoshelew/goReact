package dto

// Seat DTO
type Seat struct {
	HotelRoomSeatID int    `json:"seatId"`
	Description     string `json:"desc"`
	IsFree          bool   `json:"isFree"`
}
