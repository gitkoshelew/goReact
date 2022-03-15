package model_test

import (
	"admin/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		e       func() *model.Employee
		isValid bool
	}{
		{
			name: "valid",
			e: func() *model.Employee {
				return model.TestEmployee()
			},
			isValid: true,
		},
		{
			name: "valid user",
			e: func() *model.Employee {
				u := model.TestUser()
				e := model.TestEmployee()
				e.User = *u
				return e
			},
			isValid: true,
		},
		{
			name: "valid position",
			e: func() *model.Employee {
				e := model.TestEmployee()
				e.Position = model.OwnerPosition
				return e
			},
			isValid: true,
		},
		{
			name: "valid position",
			e: func() *model.Employee {
				e := model.TestEmployee()
				e.Position = model.EmployeePosition
				return e
			},
			isValid: true,
		},
		{
			name: "invalid position",
			e: func() *model.Employee {
				e := model.TestEmployee()
				e.Position = "wwww"
				return e
			},
			isValid: false,
		},
		{
			name: "valid hotel",
			e: func() *model.Employee {
				h := model.TestHotel()
				e := model.TestEmployee()
				e.Hotel = *h
				return e
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.e().Validate())
			} else {
				assert.Error(t, tc.e().Validate())
			}
		})
	}
}
