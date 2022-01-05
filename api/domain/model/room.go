package model

// Room struct
type Room struct {
	RoomID       int     `json:"roomId"`
	RoomNumber   int     `json:"roomNum"`
	PetType      PetType `json:"petType"`
	Hotel        Hotel
	RoomPhotoURL string `json:"roomPhotoUrl"`
}
