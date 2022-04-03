package model_test

import (
	"admin/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeat_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.Seat
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Seat {
				return model.TestSeat()
			},
			isValid: true,
		},
		{
			name: "without descriptions",
			s: func() *model.Seat {
				s := model.TestSeat()
				s.Description = ""
				return s
			},
			isValid: false,
		},
		{
			name: "valid Room",
			s: func() *model.Seat {
				r := model.TestRoom()
				s := model.TestSeat()
				s.Room = *r
				return s
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
