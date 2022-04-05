package model_test

import (
	"admin/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		e       func() *model.EmployeeDTO
		isValid bool
	}{
		{
			name: "valid",
			e: func() *model.EmployeeDTO {
				return model.TestEmployeeDTO()
			},
			isValid: true,
		},
		{
			name: "nil employee",
			e: func() *model.EmployeeDTO {
				return &model.EmployeeDTO{}
			},
			isValid: false,
		},
		{
			name: "invalid UserID",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.UserID = 0
				return employee
			},
			isValid: false,
		},
		{
			name: "invalid HotelID",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.HotelID = 0
				return employee
			},
			isValid: false,
		},
		{
			name: "invalid Position",
			e: func() *model.EmployeeDTO {
				employee := model.TestEmployeeDTO()
				employee.Position = "invalid position"
				return employee
			},
			isValid: false,
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
