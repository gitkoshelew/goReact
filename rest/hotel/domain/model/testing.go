package model

import (
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

// TestSeat ...
func TestSeat() *Seat {
	rentFrom := time.Time{}.AddDate(20220, 2, 2)
	rentTo := time.Time{}.AddDate(2022, 3, 2)
	return &Seat{
		Description: "Description of seat",
		RentFrom:    &rentFrom,
		RentTo:      &rentTo,
		Room:        *TestRoom(),
	}
}

// TestEmployee ...
func TestEmployee() *Employee {
	dateOfBirth := time.Time{}.AddDate(2000, 2, 2)
	return &Employee{
		EmployeeID:  1,
		Email:       "email@example.org",
		Role:        EmployeeRole,
		Verified:    true,
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
	return &EmployeeDTO{
		EmployeeID:  1,
		Email:       "email@example.org",
		Role:        string(EmployeeRole),
		Verified:    true,
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
