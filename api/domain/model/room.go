package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Room struct
type Room struct {
	RoomID       int     `json:"roomId"`
	RoomNumber   int     `json:"roomNum"`
	PetType      PetType `json:"petType"`
	Hotel        Hotel
	RoomPhotoURL string `json:"roomPhotoUrl"`
}

// Validate ...
func (r *Room) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.RoomNumber, validation.Required),
		validation.Field(&r.PetType, validation.Required, validation.By(IsPetType)),
		validation.Field(&r.Hotel, validation.Required),
		validation.Field(&r.RoomPhotoURL, validation.Required, validation.Length(2, 40)),
	)
}
