package reqandresp

import (
	"goReact/domain/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

// FreeSeatsSearching struct
type FreeSeatsSearching struct {
	HotelID int    `json:"hotelId"`
	PetType string `json:"petType"`
}

// Validate ...
func (fss *FreeSeatsSearching) Validate() error {
	return validation.ValidateStruct(
		fss,
		validation.Field(&fss.HotelID, validation.Required, validation.By(model.IsValidID)),
		validation.Field(&fss.PetType, validation.Required, validation.By(model.IsPetType)),
	)
}
