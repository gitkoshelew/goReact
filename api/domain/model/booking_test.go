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
		b       func() *model.BookingDTO
		isValid bool
	}{
		{
			name: "valid",
			b: func() *model.BookingDTO {
				return model.TestBooking()
			},
			isValid: true,
		},
		{
			name: "valid seat",
			b: func() *model.BookingDTO {
				b := model.TestBooking()
				s := model.TestSeat()
				p := model.TestPet()
				e := model.TestEmployee()
				b.SeatID = s.SeatID
				b.PetID = p.PetID
				b.EmployeeID = e.EmployeeID

				return b
			},
			isValid: true,
		},
		{
			name: "valid status",
			b: func() *model.BookingDTO {
				b := model.TestBooking()
				b.Status = string(model.BookingStatusCancelled)
				return b
			},
			isValid: true,
		},
		{
			name: "invalid position",
			b: func() *model.BookingDTO {
				b := model.TestBooking()
				b.Status = "st"
				return b
			},
			isValid: false,
		},
		{
			name: "valid EndDate",
			b: func() *model.BookingDTO {
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
