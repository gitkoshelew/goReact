package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionsEmployeeRepository_GetAll(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		s       func() *store.Store
		isValid bool
	}{
		{
			name: "valid",
			s: func() *store.Store {
				testStore.Open()
				return testStore
			},
			isValid: true,
		},
		{
			name: "DB closed",
			s: func() *store.Store {
				st := testStore
				st.Close()
				return st
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := tc.s().Permissions().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Permissions().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestPermissionsEmployeeRepository_SetForEmployee(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.PermissionsEmployees
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.PermissionsEmployees {
				testStore.Open()
				pe := model.TestPermissionsEmployees()
				return pe
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.PermissionsEmployees {
				testStore.Close()
				return nil
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.PermissionsEmployee().SetForEmployee(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.PermissionsEmployee().SetForEmployee(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestPermissionsEmployeeRepository__ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.PermissionsEmployeesDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.PermissionsEmployeesDTO {
				testStore.Open()

				pe := model.TestPermissionsEmployeesDTO()

				return pe
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.PermissionsEmployeesDTO {
				testStore.Close()

				pe := model.TestPermissionsEmployeesDTO()

				return pe
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.PermissionsEmployee().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.PermissionsEmployee().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
