package dto

// Room DTO
type Room struct {
	HotelRoomID int    `json:"roomId"`
	RoomNumber  int    `json:"roomNum"`
	PetType     string `json:"petType"`
	SeatsID     []int  `json:"seatsIds"`
}
