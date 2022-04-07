package reqandresp

import (
	"admin/domain/model"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// FreeSeatsSearching struct
type FreeSeatsSearching struct {
	HotelID  int        `json:"hotelId"`
	PetType  string     `json:"petType"`
	RentFrom *time.Time `json:"rentFrom"`
	RentTo   *time.Time `json:"rentTo"`
}

// Validate ...
func (fss *FreeSeatsSearching) Validate() error {
	return validation.ValidateStruct(
		fss,
		validation.Field(&fss.HotelID, validation.Required, validation.By(model.IsValidID)),
		validation.Field(&fss.PetType, validation.Required, validation.By(model.IsPetType)),
		validation.Field(&fss.RentFrom, validation.Required, validation.By(model.IsValidStartDate)),
		validation.Field(&fss.RentTo, validation.Required, validation.By(model.IsValidEndDate)),
	)
}
