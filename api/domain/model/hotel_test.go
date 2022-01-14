package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotel_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		h       func() *model.Hotel
		isValid bool
	}{{
		name: "valid",
		h: func() *model.Hotel {
			return model.TestHotel(t)
		},
		isValid: true,
		},
		{
			name: "Invalid Name",
			h: func() *model.Hotel {
				h := model.TestHotel(t)
				h.Name = "Name@123"
				return h
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			h: func() *model.Hotel {
				h := model.TestHotel(t)
				h.Name = ""
				return h
			},
			isValid: false,
		},
		{
			name: "Empty Address",
			h: func() *model.Hotel {
				h := model.TestHotel(t)
				h.Address = ""
				return h
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.h().Validate())
			} else {
				assert.Error(t, tc.h().Validate())
			}
		})
	}
}
