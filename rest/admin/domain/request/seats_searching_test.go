package request

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPet_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		req     func() *FreeSeatsSearching
		isValid bool
	}{
		{
			name: "valid",
			req: func() *FreeSeatsSearching {
				return TestFreeSeatsSearching()
			},
			isValid: true,
		},
		{
			name: "invalid HotelID",
			req: func() *FreeSeatsSearching {
				request := TestFreeSeatsSearching()
				request.HotelID = 0
				return request
			},
			isValid: false,
		},
		{
			name: "invalid PetType",
			req: func() *FreeSeatsSearching {
				request := TestFreeSeatsSearching()
				request.PetType = "invalid"
				return request
			},
			isValid: false,
		},
		{
			name: "invalid SratDate",
			req: func() *FreeSeatsSearching {
				request := TestFreeSeatsSearching()
				rentFrom := time.Now().AddDate(0, 0, -1)
				request.RentFrom = &rentFrom
				return request
			},
			isValid: false,
		},
		{
			name: "invalid EndDate",
			req: func() *FreeSeatsSearching {
				request := TestFreeSeatsSearching()
				RentTo := time.Now().AddDate(0, 0, -1)
				request.RentFrom = &RentTo
				return request
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.req().Validate())
			} else {
				assert.Error(t, tc.req().Validate())
			}
		})
	}
}
