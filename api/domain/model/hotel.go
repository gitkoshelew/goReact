package model

import validation "github.com/go-ozzo/ozzo-validation"

// Hotel struct
type Hotel struct {
	HotelID     int       `json:"hotelId"`
	Name        string    `json:"nameId"`
	Address     string    `json:"addressId"`
	Coordinates []float64 `json:"coordinates"`
}

// Validate ...
func (h *Hotel) Validate() error {
	return validation.ValidateStruct(
		h,
		validation.Field(&h.Name, validation.Required, validation.By(IsSQL), validation.Length(1, 20)),
		validation.Field(&h.Address, validation.Required, validation.By(IsSQL), validation.Length(10, 40)),
		validation.Field(&h.Coordinates, validation.Required, validation.Each(validation.Required)),
	)
}
