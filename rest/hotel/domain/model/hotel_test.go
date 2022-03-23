package model_test

import (
	"hotel/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotel_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		model   func() *model.Hotel
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Hotel {
				return model.TestHotel()
			},
			isValid: true,
		},
		{
			name: "Empty Name",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Name = ""
				return hotel
			},
			isValid: false,
		},
		{
			name: "SQL Name",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Name = "ALtE  *)/8 R"
				return hotel
			},
			isValid: false,
		},
		{
			name: "Long Name",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Name = "1234567891011121314151617181920"
				return hotel
			},
			isValid: false,
		},
		{
			name: "Empty Address",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Address = ""
				return hotel
			},
			isValid: false,
		},
		{
			name: "SQL Address",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Address = "ALt E  *)/8 R"
				return hotel
			},
			isValid: false,
		},
		{
			name: "short Address",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Address = "12345"
				return hotel
			},
			isValid: false,
		},
		{
			name: "Long Address",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Address = "123456789101112131415161718192021222324252627282930313234353637383940"
				return hotel
			},
			isValid: false,
		},
		{
			name: "empty Coordinates",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Coordinates = []float64{}
				return hotel
			},
			isValid: false,
		},
		{
			name: "nil Coordinates",
			model: func() *model.Hotel {
				hotel := model.TestHotel()
				hotel.Coordinates = nil
				return hotel
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.model().Validate())
			} else {
				assert.Error(t, tc.model().Validate())
			}
		})
	}
}
