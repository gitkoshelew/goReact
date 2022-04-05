package model_test

import (
	"goReact/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//TestBooking_Validate..
func TestBooking_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		b       func() *model.BookingDTO
		isValid bool
	}{
		{
			name: "valid",
			b: func() *model.BookingDTO {
				return model.TestBookingDTO()
			},
			isValid: true,
		},
		{
			name: "invalid seatID",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.SeatID = 0
				return b
			},
			isValid: false,
		}, {
			name: "invalid EmployeeID",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.EmployeeID = 0
				return b
			},
			isValid: false,
		},
		{
			name: "invalid PetID",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.PetID = 0
				return b
			},
			isValid: false,
		},
		{
			name: "invalid position",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.Status = "st"
				return b
			},
			isValid: false,
		},
		{
			name: "invalid status",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.Status = "invalid status"
				return b
			},
			isValid: false,
		},
		{
			name: "Empty status",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.Status = ""
				return b
			},
			isValid: false,
		},
		{
			name: "invalid EndDate",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				endDate := time.Now().AddDate(0, 0, -10)
				b.EndDate = &endDate
				return b
			},
			isValid: false,
		},
		{
			name: "invalid StartDate",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				startDate := time.Now().AddDate(0, 0, -10)
				b.StartDate = &startDate
				return b
			},
			isValid: false,
		},
		{
			name: "Nil StartDate",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.StartDate = nil
				return b
			},
			isValid: false,
		},
		{
			name: "Nil EndDate",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.EndDate = nil
				return b
			},
			isValid: false,
		},
		{
			name: "SQL notes",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.Notes = "alTe&5 *&2 99*6 &89 --) R"
				return b
			},
			isValid: false,
		},
		{
			name: "Nil paid",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.Paid = nil
				return b
			},
			isValid: false,
		},
		{
			name: "Invalid transactionID",
			b: func() *model.BookingDTO {
				b := model.TestBookingDTO()
				b.TransactionID = 0
				return b
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
