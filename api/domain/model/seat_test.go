package model_test

import (
	"goReact/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSeat_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		b       func() *model.SeatDTO
		isValid bool
	}{
		{
			name: "valid",
			b: func() *model.SeatDTO {
				return model.TestSeatDTO()
			},
			isValid: true,
		},
		{
			name: "invalid RoomID",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				seat.RoomID = -1
				return seat
			},
			isValid: false,
		},
		{
			name: "invalid RentFrom",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				rentFrom := time.Now().AddDate(0, 0, -10)
				seat.RentFrom = &rentFrom
				return seat
			},
			isValid: false,
		},
		{
			name: "invalid RentFrom",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				seat.RentFrom = nil
				return seat
			},
			isValid: false,
		},
		{
			name: "invalid RentTo",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				rentFrom := time.Now().AddDate(0, 0, -10)
				seat.RentFrom = &rentFrom
				return seat
			},
			isValid: false,
		},
		{
			name: "nil RentTo",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				seat.RentTo = nil
				return seat
			},
			isValid: false,
		},
		{
			name: "invalid Price",
			b: func() *model.SeatDTO {
				seat := model.TestSeatDTO()
				seat.Price = -10
				return seat
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.b().Validate())
			} else {
				assert.Error(t, tc.b().Validate())
			}
		})
	}
}
