package model_test

import (
	"goReact/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBooking_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		b       func() *model.Booking
		isValid bool
	}{
		{
			name: "valid",
			b: func() *model.Booking {
				return model.TestBooking()
			},
			isValid: true,
		},
		{
			name: "valid seat",
			b: func() *model.Booking {
				b := model.TestBooking()
				s := model.TestSeat()
				p := model.TestPet()
				e := model.TestEmployee()
				b.Seat = *s
				b.Pet = *p
				b.Employee = *e

				return b
			},
			isValid: true,
		},
		{
			name: "valid status",
			b: func() *model.Booking {
				b := model.TestBooking()
				b.Status = model.BookingStatusCancelled
				return b
			},
			isValid: true,
		},
		{
			name: "invalid position",
			b: func() *model.Booking {
				b := model.TestBooking()
				b.Status = "st"
				return b
			},
			isValid: false,
		},
		{
			name: "valid EndDate",
			b: func() *model.Booking {
				a := time.Time{}.AddDate(2000, 2, 2)
				b := model.TestBooking()
				b.EndDate = &a
				return b
			},
			isValid: true,
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
