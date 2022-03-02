package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Hotel struct
type Hotel struct {
	HotelID     int        `json:"hotelId"`
	Name        string     `json:"nameId"`
	Address     string     `json:"addressId"`
	Coordinates [2]float64 `json:"coordinates"` // coordinates : lat , lon
}

// Validate ...
func (h *Hotel) Validate() error {
	return validation.ValidateStruct(
		h,
		validation.Field(&h.Name, validation.Required, validation.Length(1, 20)),
		validation.Field(&h.Address, validation.Required, validation.Length(10, 40)),
		validation.Field(&h.Coordinates, validation.Each(validation.Required)),
	)
}
