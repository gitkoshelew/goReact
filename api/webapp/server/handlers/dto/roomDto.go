package dto

// RoomDto ...
type RoomDto struct {
	RoomID     int    `json:"roomId"`
	RoomNumber int    `json:"roomNum"`
	PetType    string `json:"petType"`
	HotelID    int    `json:"hotelId"`
}
