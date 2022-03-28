package store_test

import (
	"admin/domain/store"
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

	type id struct {
		permissionID int
		employeeId   int
	}
	testCases := []struct {
		name string
		id id
		isValid bool
	}{
		{
			name: "valid",
			id: func() id {
				testStore.Open()
				permissionID := id.Permission
				employeeID := id.Employee
				var id struct{} = id{
					pepermissionID: pepermissionID,
					eemployeeId:    employeeID,
				}
				return id
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
				result, err := testStore.Permissions().SetForEmployee()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Permissions().SetForEmployee()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
