package request

import (
	"hotel/domain/model"
	"time"
)

// TestFreeSeatsSearching ...
func TestFreeSeatsSearching() *FreeSeatsSearching {
	rentFrom := time.Now().AddDate(0, 0, 1)
	rentTo := time.Now().AddDate(0, 0, 10)

	return &FreeSeatsSearching{
		HotelID:  1,
		PetType:  string(model.PetTypeCat),
		RentFrom: &rentFrom,
		RentTo:   &rentTo,
	}
}
