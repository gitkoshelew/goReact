package reqandresp

import (
	"goReact/domain/model"
)

// TestFreeSeatsSearching ...
func TestFreeSeatsSearching() *FreeSeatsSearching {

	return &FreeSeatsSearching{
		HotelID: 1,
		PetType: string(model.PetTypeCat),
	}
}
