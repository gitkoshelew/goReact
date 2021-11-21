package entity

type HotelId int

type Hotel struct {
	Id       HotelId
	Name     string
	Rooms    []HotelRoom
	Bookings []Booking
}

func (hotel *Hotel) RoomsCount() int {
	if hotel == nil {
		return 0
	}
	return len(hotel.Rooms)
}

func (hotel *Hotel) SeatsCount() int {
	count := 0
	for _, room := range hotel.Rooms {
		count += room.SeatsCount()
	}
	return count
}
