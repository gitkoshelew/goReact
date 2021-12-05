package entity

// HotelRoom struct
type HotelRoom struct {
	HotelRoomID int             `json:"roomId"`
	RoomNumber  int             `json:"roomNum"`
	PetType     PetType         `json:"petType"`
	Seats       []HotelRoomSeat `json:"seatsIds"`
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

// GetRoomByID returns Room by id from storage
func GetRoomByID(id int) HotelRoom {
	var room HotelRoom
	for _, r := range GetHotelRooms() {
		if id == r.HotelRoomID {
			room = r
		}
	}
	return room
}

// GetRoomsByID returns []HotelRoom by []ids from storage
func GetRoomsByID(ids []int) []HotelRoom {
	var rooms []HotelRoom
	for _, id := range ids {
		rooms = append(rooms, GetRoomByID(id))
	}
	return rooms
}
