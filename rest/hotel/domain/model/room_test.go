package model_test

import (
	"hotel/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoom_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		r       func() *model.RoomDTO
		isValid bool
	}{
		{
			name: "valid",
			r: func() *model.RoomDTO {
				return model.TestRoomDTO()
			},
			isValid: true,
		},
		{
			name: "invalid pet type",
			r: func() *model.RoomDTO {
				r := model.TestRoomDTO()
				r.PetType = "invalid"
				return r
			},
			isValid: false,
		},
		{
			name: "invalid Room Number",
			r: func() *model.RoomDTO {
				r := model.TestRoomDTO()
				r.RoomNumber = -5
				return r
			},
			isValid: false,
		},
		{
			name: "Large Room Number",
			r: func() *model.RoomDTO {
				r := model.TestRoomDTO()
				r.RoomNumber = 99999999999
				return r
			},
			isValid: false,
		},
		{
			name: "Invalid HotelID",
			r: func() *model.RoomDTO {
				r := model.TestRoomDTO()
				r.HotelID = -1
				return r
			},
			isValid: false,
		},
		{
			name: "SQL PhotoURL",
			r: func() *model.RoomDTO {
				r := model.TestRoomDTO()
				r.PhotoURL = "Al t^&*745Er"
				return r
			},
			isValid: false,
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
