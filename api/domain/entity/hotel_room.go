package entity

// HotelRoom struct
type HotelRoom struct {
	HotelRoomID int
	RoomNumber  int
	PetType     PetType
	Seats       []HotelRoomSeat
}

// SetRoomNumber sets HotelRooms number
func (h *HotelRoom) SetRoomNumber(i int) {
	h.RoomNumber = i
}

// SetPetType sets HotelRooms pet type
func (h *HotelRoom) SetPetType(p PetType) {
	h.PetType = p
}

// SetSeats sets HotelRooms seats
func (h *HotelRoom) SetSeats(s []HotelRoomSeat) {
	h.Seats = s
}

// SeatsCount returns a numbers of Seats in Hotel Room
func (h *HotelRoom) SeatsCount() int {
	if h == nil {
		return 0
	}
	return len(h.Seats)
}
