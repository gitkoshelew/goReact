package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Room struct
type Room struct {
	RoomID     int     `json:"roomId"`
	RoomNumber int     `json:"roomNum"`
	PetType    PetType `json:"petType"`
	Hotel      Hotel
	PhotoURL   string `json:"photoUrl,omitempty"`
}

// RoomDTO struct
type RoomDTO struct {
	RoomID     int    `json:"roomId"`
	RoomNumber int    `json:"roomNum"`
	PetType    string `json:"petType"`
	HotelID    int    `json:"hotelId"`
	PhotoURL   string `json:"photoUrl,omitempty"`
}

// Validate ...
func (r *RoomDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.RoomNumber, validation.Required, validation.Min(1), validation.Max(999)),
		validation.Field(&r.PetType, validation.Required, validation.By(IsPetType)),
		validation.Field(&r.HotelID, validation.Required, validation.By(IsValidID)),
		validation.Field(&r.PhotoURL, validation.By(IsSQL)),
	)
}
