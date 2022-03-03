package hotel

import "time"

// HotelDTO ...
type HotelDTO struct {
	HotelID     int    `json:"hotelId"`
	Name        string `json:"nameId"`
	Address     string `json:"addressId"`
	Coordinates string `json:"coordinates"` // coordinates : lat , lon
}

// PetType ...
type PetType string

const (
	// PetTypeCat ...
	PetTypeCat PetType = "cat"
	// PetTypeDog ...
	PetTypeDog PetType = "dog"
)

// RoomDTO ...
type RoomDTO struct {
	RoomID       int     `json:"roomId"`
	RoomNumber   int     `json:"roomNum"`
	PetType      PetType `json:"petType"`
	HotelID      int     `json:"hotelID"`
	RoomPhotoURL string  `json:"roomPhotoUrl"`
}

// SeatDTO ...
type SeatDTO struct {
	SeatID      int       `json:"seatId"`
	RoomID      int       `json:"roomId"`
	Description string    `json:"description,omitempty"`
	RentFrom    time.Time `json:"rentFrom"`
	RentTo      time.Time `json:"rentTo"`
}

// Position ...
type Position string

// Position constants
const (
	ManagerPosition  Position = "manager"
	EmployeePosition Position = "employee"
	OwnerPosition    Position = "owner"
	AdminPosition    Position = "admin"
)

// EmployeeDTO ...
type EmployeeDTO struct {
	EmployeeID int      `json:"employeeId"`
	UserID     int      `json:"userId"`
	HotelID    int      `json:"hotelId"`
	Position   Position `json:"position"`
}
