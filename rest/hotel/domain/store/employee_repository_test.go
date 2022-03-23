package store_test

import (
	"hotel/domain/model"
	"hotel/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Employee
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Employee {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Hotel.HotelID = id.Hotel

				return employee
			},
			isValid: true,
		},
		{
			name: "email in use",
			model: func() *model.Employee {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Email = "new@mail.org"
				testStore.Employee().Create(employee)

				return employee
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Employee {
				testStore.Close()
				return model.TestEmployee()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Employee().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Employee().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestEmployeeRepository_GetAll(t *testing.T) {
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
				result, err := tc.s().Employee().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Employee().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestEmployeeRepository_FindByID(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()
				return id.Employee
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.Employee
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Employee().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Employee().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestEmployeeRepository_Delete(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Hotel.HotelID = id.Hotel
				id, _ := testStore.Employee().Create(employee)

				return *id
			},
			isValid: true,
		},
		{
			name: "Invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.Employee
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Employee().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Employee().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		email   func() string
		isValid bool
	}{
		{
			name: "valid",
			email: func() string {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Hotel.HotelID = id.Hotel
				employee.Email = "searching1@email.org"
				testStore.Employee().Create(employee)

				return employee.Email
			},
			isValid: true,
		},
		{
			name: "invalid Email",
			email: func() string {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Hotel.HotelID = id.Hotel
				employee.Email = "searching2@email.org"
				testStore.Employee().Create(employee)

				return "notThis@email.org"
			},
			isValid: false,
		},
		{
			name: "DB closed",
			email: func() string {
				testStore.Close()

				employee := model.TestEmployee()
				employee.Hotel.HotelID = id.Hotel
				employee.Email = "searching3@email.org"
				testStore.Employee().Create(employee)

				return employee.Email
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Employee().FindByEmail(tc.email())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Employee().FindByEmail(tc.email())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestEmployeeRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Employee
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Employee {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Email = "new@mail.org"
				employee.EmployeeID = id.Employee
				employee.Hotel.HotelID = id.Hotel

				return employee
			},
			isValid: true,
		},
		{
			name: "Email is used",
			model: func() *model.Employee {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Email = "inUse@mail.org"
				employee.EmployeeID = id.Employee
				employee.Hotel.HotelID = id.Hotel
				testStore.Employee().Create(employee)
				return employee
			},
			isValid: false,
		},
		{
			name: "invalid ID",
			model: func() *model.Employee {
				testStore.Open()

				employee := model.TestEmployee()
				employee.EmployeeID = 0
				employee.Hotel.HotelID = id.Hotel

				return employee
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Employee {
				testStore.Close()

				employee := model.TestEmployee()
				employee.EmployeeID = id.Employee
				employee.Hotel.HotelID = id.Hotel

				return employee
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Employee().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Employee().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestEmployeeRepository_ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.EmployeeDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.EmployeeDTO {
				testStore.Open()

				employee := model.TestEmployeeDTO()
				employee.HotelID = id.Hotel

				return employee
			},
			isValid: true,
		},
		{
			name: "Invalid HotelID",
			model: func() *model.EmployeeDTO {
				testStore.Open()

				employee := model.TestEmployeeDTO()
				employee.HotelID = 0

				return employee
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.EmployeeDTO {
				testStore.Close()

				employee := model.TestEmployeeDTO()
				employee.HotelID = id.Hotel

				return employee
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Employee().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Employee().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserRepository_EmailCheck(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		email   func() string
		isValid bool
		isExist bool
	}{
		{
			name: "Email in use",
			email: func() string {
				testStore.Open()

				employee := model.TestEmployee()
				employee.Email = "inusing@email.org"
				employee.Hotel.HotelID = id.Hotel
				testStore.Employee().Create(employee)

				return employee.Email
			},
			isValid: true,
			isExist: true,
		},
		{
			name: "Email is not in use",
			email: func() string {
				testStore.Open()

				return "unusing@email.org"
			},
			isValid: true,
			isExist: false,
		},
		{
			name: "DB closed",
			email: func() string {
				testStore.Close()

				return "unusing2@email.org"
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Employee().EmailCheck(tc.email())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if tc.isExist {
					assert.True(t, *result)
				} else {
					assert.False(t, *result)
				}

			} else {
				result, err := testStore.Employee().EmailCheck(tc.email())
				testStore.Close()
				assert.Error(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
