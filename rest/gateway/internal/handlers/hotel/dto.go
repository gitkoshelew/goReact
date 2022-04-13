package hotel

import "time"

// HotelDTO ...
type HotelDTO struct {
	HotelID     int       `json:"hotelId"`
	Name        string    `json:"nameId"`
	Address     string    `json:"addressId"`
	Coordinates []float64 `json:"coordinates"` // coordinates : lat , lon
}

// PetType ...
type PetType string

const (
	// PetTypeCat ...
	PetTypeCat PetType = "cat"
	// PetTypeDog ...
	PetTypeDog PetType = "dog"
)

// RoomDTO struct
type RoomDTO struct {
	RoomID      int     `json:"roomId"`
	RoomNumber  int     `json:"roomNum"`
	PetType     string  `json:"petType"`
	HotelID     int     `json:"hotelId"`
	PhotoURL    string  `json:"photoUrl,omitempty"`
	Description string  `json:"description,omitempty"`
	Square      float64 `json:"square,omitempty"`
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
