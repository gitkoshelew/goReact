package model

import (
	"hotel/pkg/pagination"
	"time"
)

// TestHotel instance of hotel
func TestHotel() *Hotel {
	return &Hotel{
		HotelID:     1,
		Name:        "Name",
		Address:     "Minsk ul sovetskaya 18",
		Coordinates: []float64{53.89909164468815, 27.498996594142426},
	}
}

// TestHotelDTO ...
func TestHotelDTO() *Hotel {
	return &Hotel{
		HotelID:     1,
		Name:        "Name",
		Address:     "Minsk ul sovetskaya 18",
		Coordinates: []float64{53.89909164468815, 27.498996594142426},
	}
}

// TestRoom ...
func TestRoom() *Room {
	return &Room{
		RoomID:     1,
		RoomNumber: 1,
		PetType:    PetTypeCat,
		Hotel:      *TestHotel(),
		PhotoURL:   "/photo/1",
	}
}

// TestRoomDTO ...
func TestRoomDTO() *RoomDTO {
	return &RoomDTO{
		RoomID:     1,
		RoomNumber: 1,
		PetType:    "cat",
		HotelID:    TestHotel().HotelID,
		PhotoURL:   "/photo/1",
	}
}

// TestSeat ...
func TestSeat() *Seat {
	rentFrom := time.Now().AddDate(0, 0, 1)
	rentTo := time.Now().AddDate(0, 0, 10)
	return &Seat{
		SeatID:      1,
		Description: "Description of seat",
		RentFrom:    &rentFrom,
		RentTo:      &rentTo,
		Room:        *TestRoom(),
	}
}

// TestSeatDTO ...
func TestSeatDTO() *SeatDTO {
	rentFrom := time.Now().AddDate(0, 0, 1)
	rentTo := time.Now().AddDate(0, 0, 10)
	return &SeatDTO{
		SeatID:      1,
		Description: "Description of seat",
		RentFrom:    &rentFrom,
		RentTo:      &rentTo,
		RoomID:      1,
	}
}

// TestEmployee ...
func TestEmployee() *Employee {
	dateOfBirth := time.Time{}.AddDate(2000, 2, 2)
	verified := true
	return &Employee{
		EmployeeID:  1,
		Email:       "email@example.org",
		Role:        EmployeeRole,
		Verified:    &verified,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         SexMale,
		DateOfBirth: &dateOfBirth,
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
		Hotel:       *TestHotel(),
		Position:    OwnerPosition,
	}
}

// TestEmployeeDTO ...
func TestEmployeeDTO() *EmployeeDTO {
	dateOfBirth := time.Time{}.AddDate(2000, 2, 2)
	verified := true
	return &EmployeeDTO{
		EmployeeID:  1,
		Email:       "email@example.org",
		Role:        string(EmployeeRole),
		Verified:    &verified,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         string(SexMale),
		DateOfBirth: &dateOfBirth,
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
		HotelID:     1,
		Position:    string(OwnerPosition),
	}
}

// TestPage ...
func TestPage() *pagination.Page {
	return &pagination.Page{
		PageNumber: 1,
		PageSize:   10,
	}
}
