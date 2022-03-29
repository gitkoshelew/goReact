package model_test

import (
	"admin/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionsEmployee_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		pe      func() *model.PermissionsEmployeesDTO
		isValid bool
	}{
		{name: "valid",
			pe: func() *model.PermissionsEmployeesDTO {
				return model.TestPermissionsEmployeesDTO()
			},
			isValid: true,
		},
		{name: "invalid employee id",
			pe: func() *model.PermissionsEmployeesDTO {
				e := model.TestPermissionsEmployeesDTO()
				e.EmployeeID = 0
				return e
			},
			isValid: false,
		},
		{name: "invalid permission id",
			pe: func() *model.PermissionsEmployeesDTO {
				e := model.TestPermissionsEmployeesDTO()
				e.PermissionsID = 0
				return e
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.pe().Validate())
			} else {
				assert.Error(t, tc.pe().Validate())
			}
		})
	}
}
