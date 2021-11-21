package entity

type HotelRoomId int

type HotelRoom struct {
	Id         HotelRoomId
	RoomNumber string
	PetType    PetType
	Seats      []HotelRoomSeat
}

func (room *HotelRoom) SeatsCount() int {
	if room == nil {
		return 0
	}
	return len(room.Seats)
}
