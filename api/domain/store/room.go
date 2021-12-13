package store

// Room struct
type Room struct {
	RoomID     int     `json:"roomId"`
	RoomNumber int     `json:"roomNum"`
	PetType    PetType `json:"petType"`
	Hotel      Hotel
}

// SetRoomNumber sets Rooms number
func (r *Room) SetRoomNumber(i int) {
	r.RoomNumber = i
}

// SetPetType sets Rooms pet type
func (r *Room) SetPetType(p PetType) {
	r.PetType = p
}

// SetHotel sets Rooms hotel
func (r *Room) SetHotel(h Hotel) {
	r.Hotel = h
}
