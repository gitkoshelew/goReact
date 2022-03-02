package model

import "time"

// Testing instance of hotel
func TestHotel() *Hotel {
	return &Hotel{
		HotelID: 1,
		Name:    "Name",
		Address: "Minsk ul sovetskaya 18",
		Coordinates: [2]float64 {53.89909164468815, 27.498996594142426},
	}
}

func TestRoom() *Room {
	return &Room{
		RoomID:       1,
		RoomNumber:   1,
		PetType:      PetTypeCat,
		Hotel:        *TestHotel(),
		RoomPhotoURL: "/photo/1",
	}
}

func TestSeat() *Seat {
	return &Seat{
		Description: "Description of seat",
		RentFrom:    time.Time{}.AddDate(20220, 2, 2),
		RentTo:      time.Time{}.AddDate(2022, 3, 2),
		Room:        *TestRoom(),
	}
}

func TestEmployee() *Employee {
	return &Employee{
		EmployeeID: 1,
		UserID:   1,
		Hotel:    *TestHotel(),
		Position: OwnerPosition,
	}
}
