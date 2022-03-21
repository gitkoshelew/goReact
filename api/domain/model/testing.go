package model

import (
	"goReact/webapp/server/handler/pagination"
	"time"
)

// TestUser ...
func TestUser() *User {
	return &User{
		UserID:      1,
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

// TestEmployee ...
func TestEmployee() *Employee {
	return &Employee{
		EmployeeID: 1,
		User:       *TestUser(),
		Hotel:      *TestHotel(),
		Position:   OwnerPosition,
	}
}

// TestEmployeeDTO ...
func TestEmployeeDTO() *EmployeeDTO {
	return &EmployeeDTO{
		EmployeeID: 1,
		UserID:     1,
		HotelID:    1,
		Position:   string(OwnerPosition),
	}
}

// TestPet ...
func TestPet() *Pet {
	return &Pet{
		PetID:    1,
		Name:     "Name",
		Type:     PetTypeCat,
		Weight:   1,
		Diseases: "Diseases",
		Owner:    *TestUser(),
		PhotoURL: "/",
	}
}

// TestPetDTO ...
func TestPetDTO() *PetDTO {
	return &PetDTO{
		PetID:    1,
		Name:     "Name",
		Type:     string(PetTypeCat),
		Weight:   1,
		Diseases: "Diseases",
		OwnerID:  1,
		PhotoURL: "/",
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

// TestBookingDTO ...
func TestBookingDTO() *BookingDTO {
	a := time.Now().AddDate(0, 0, 1)
	b := time.Now().AddDate(0, 0, 10)
	paid := true
	return &BookingDTO{
		SeatID:        TestSeat().SeatID,
		PetID:         TestPet().PetID,
		EmployeeID:    TestEmployee().EmployeeID,
		Status:        string(BookingStatusInProgress),
		StartDate:     &a,
		EndDate:       &b,
		Notes:         "Notes",
		TransactionID: 1,
		Paid:          &paid,
	}
}

// TestBooking ...
func TestBooking() *Booking {
	a := time.Now().AddDate(0, 0, 1)
	b := time.Now().AddDate(0, 0, 10)
	paid := true
	return &Booking{
		Seat:          *TestSeat(),
		Pet:           *TestPet(),
		Employee:      *TestEmployee(),
		Status:        BookingStatusInProgress,
		StartDate:     &a,
		EndDate:       &b,
		Notes:         "Notes",
		TransactionID: 1,
		Paid:          &paid,
	}
}
