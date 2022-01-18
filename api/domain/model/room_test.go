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
				return model.TestRoom()
			},
			isValid: true,
		},
		{
			name: "0 or less room number",
			r: func() *model.Room {
				r := model.TestRoom()
				r.RoomNumber = -5 
				return r
			},
			isValid: false,
		},
		{
			name: "invalid pet type",
			r: func() *model.Room {
				r := model.TestRoom()
				r.PetType = "invalid"
				return r
			},
			isValid: false,
		},
		{
			name: "valid pet type",
			r: func() *model.Room {
				r := model.TestRoom()
				r.PetType = model.PetTypeCat
				return r
			},
			isValid: true,
		},
		{
			name: "valid pet type",
			r: func() *model.Room {
				r := model.TestRoom()
				r.PetType = model.PetTypeDog
				return r
			},
			isValid: true,
		},
		{
			name: "valid hotel",
			r: func() *model.Room {
				h := model.TestHotel()
				r := model.TestRoom()
				r.Hotel = *h
				return r
			},
			isValid: true,
		},
<<<<<<< HEAD
=======
		{
			name: "Empty RoomPhotoURL",
			r: func() *model.Room {
				r := model.TestRoom()
				r.RoomPhotoURL = ""
				return r
			},
			isValid: false,
		},
		{
			name: "Invalid RoomPhotoURL",
			r: func() *model.Room {
				r := model.TestRoom()
				r.RoomPhotoURL = "/"
				return r
			},
			isValid: false,
		},
>>>>>>> test-hotel/room
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
