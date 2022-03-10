package model

import (
	"goReact/webapp/server/handler/pagination"
	"time"
)

// TestUser ...
func TestUser() *UserDTO {
	return &UserDTO{
		Email:       "email@example.org",
		Password:    "password",
		Role:        "client",
		Verified:    true,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         "male",
		DateOfBirth: time.Time{}.AddDate(2000, 2, 2),
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
	}

}

// TestHotel instance of hotel
func TestHotel() *Hotel {
	return &Hotel{
		HotelID:     1,
		Name:        "Name",
		Address:     "Minsk ul sovetskaya 18",
		Coordinates: "53.89909164468815, 27.498996594142426",
	}
}

// TestRoom ...
func TestRoom() *Room {
	return &Room{
		RoomID:       1,
		RoomNumber:   1,
		PetType:      PetTypeCat,
		Hotel:        *TestHotel(),
		RoomPhotoURL: "/photo/1",
	}
}

// TestRoomDTO ...
func TestRoomDTO() *RoomDTO {
	return &RoomDTO{
		RoomID:       1,
		RoomNumber:   1,
		PetType:      PetTypeCat,
		HotelID:      TestHotel().HotelID,
		RoomPhotoURL: "/photo/1",
	}
}

// TestEmployee ...
func TestEmployee() *Employee {
	return &Employee{
		EmployeeID: 1,
		User:       *TestUser().ModelFromDTO(),
		Hotel:      *TestHotel(),
		Position:   OwnerPosition,
	}
}

// TestPet ...
func TestPet() *Pet {
	return &Pet{
		Name:        "Name",
		Type:        PetTypeCat,
		Weight:      1,
		Diesieses:   "Disease",
		Owner:       *TestUser().ModelFromDTO(),
		PetPhotoURL: "/",
	}
}

// TestPage ...
func TestPage() *pagination.Page {
	return &pagination.Page{
		PageNumber: 1,
		PageSize:   10,
	}
}

// TestSeat ...
func TestSeat() *Seat {
	return &Seat{
		Description: "Description of seat",
		RentFrom:    time.Time{}.AddDate(2000, 2, 2),
		RentTo:      time.Time{}.AddDate(2001, 2, 2),
		Room:        *TestRoom(),
	}
}

// TestBooking ...
func TestBooking() *Booking {
	a := time.Time{}.AddDate(2000, 2, 2)
	b := time.Time{}.AddDate(2000, 22, 2)
	return &Booking{
		Seat:      *TestSeat(),
		Pet:       *TestPet(),
		Employee:  *TestEmployee(),
		Status:    BookingStatusInProgress,
		StartDate: &a,
		EndDate:   &b,
		Notes:     "Notes",
	}
}
