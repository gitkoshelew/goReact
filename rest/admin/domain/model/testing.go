package model

import (
	"time"
)

// TestUser ...
func TestUser() *User {
	return &User{
		Email:       "email@example.org",
		Password:    "password",
		Role:        EmployeeRole,
		Verified:    true,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         SexMale,
		DateOfBirth: time.Time{}.AddDate(2000, 2, 2),
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
	}

}

// Testing instance of hotel
func TestHotel() *Hotel {
	return &Hotel{
		HotelID:     1,
		Name:        "Name",
		Address:     "Minsk ul sovetskaya 18",
		Coordinates: "53.89909164468815, 27.498996594142426",
	}
}

// Testing instance of room
func TestRoom() *Room {
	return &Room{
		RoomID:       1,
		RoomNumber:   1,
		PetType:      PetTypeCat,
		Hotel:        *TestHotel(),
		RoomPhotoURL: "/photo/1",
	}
}

// Testing instance of employee
func TestEmployee() *Employee {
	return &Employee{
		User:     *TestUser(),
		Hotel:    *TestHotel(),
		Position: OwnerPosition,
	}
}

// Testing instance of pet
func TestPet() *Pet {
	return &Pet{
		Name:        "Name",
		Type:        PetTypeCat,
		Weight:      1,
		Diseases:    "Disease",
		Owner:       *TestUser(),
		PetPhotoURL: "/",
	}
}

// Testing instance of seat
func TestSeat() *Seat {
	return &Seat{
		Description: "Description of seat",
		RentFrom:    time.Time{}.AddDate(20220, 2, 2),
		RentTo:      time.Time{}.AddDate(2022, 3, 2),
		Room:        *TestRoom(),
	}
}

// Testing instance of booking
func TestBooking() *Booking {
	return &Booking{
		Seat:      *TestSeat(),
		Pet:       *TestPet(),
		Employee:  *TestEmployee(),
		Status:    BookingStatusInProgress,
		StartDate: time.Time{}.AddDate(2000, 2, 2),
		EndDate:   time.Time{}.AddDate(2000, 22, 2),
		Paid:      true,
		Notes:     "Notes",
	}
}
