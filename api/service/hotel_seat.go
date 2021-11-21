package service

import "goReact/domain/entity"

func GetHotelSeats() []entity.HotelRoomSeat {
	var allSeats []entity.HotelRoomSeat

	for _, room := range hotel.Rooms {
		allSeats = append(allSeats, room.Seats...)
	}

	return allSeats
}
