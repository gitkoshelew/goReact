package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoom_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		r       func() *model.Room
		isValid bool
	}{
		{
			name: "valid",
			r: func() *model.Room {
				return model.TestRoom(t)
			},
			isValid: true,
		},
		{
			name: "empty room number",
			r: func() *model.Room {
				r := model.TestRoom(t)
				r.RoomNumber = 0 ///todo
				return r
			},
			isValid: false,
		},
		{
			name: "invalid pet type",
			r: func() *model.Room {
				r := model.TestRoom(t)
				r.PetType = "invalid"
				return r
			},
			isValid: false,
		},
		{
			name: "valid pet type",
			r: func() *model.Room {
				r := model.TestRoom(t)
				r.PetType = model.PetTypeCat
				return r
			},
			isValid: true,
		},
		{
			name: "valid pet type",
			r: func() *model.Room {
				r := model.TestRoom(t)
				r.PetType = model.PetTypeDog
				return r
			},
			isValid: true,
		},
		{
			name: "valid hotel",
			r: func() *model.Room {
				h := model.TestHotel(t)
				r := model.TestRoom(t)
				r.Hotel = *h
				return r
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.r().Validate())
			} else {
				assert.Error(t, tc.r().Validate())
			}
		})
	}
}
