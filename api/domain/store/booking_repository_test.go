package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testStore, teardown = store.TestStore(&testing.T{}, host, dbName, user, password, port, sslMode)
)

func TestBookingRepository_Create(t *testing.T) {
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Booking
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Booking {
				testStore.Open()

				booking := model.TestBooking()
				booking.Employee.EmployeeID = id.Employee
				booking.Pet.PetID = id.Pet
				booking.Seat.SeatID = id.Seat

				return booking
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Booking {
				testStore.Close()
				return model.TestBooking()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Booking().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Booking().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestBookingRepository_GetAll(t *testing.T) {
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
				result, err := tc.s().Booking().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Booking().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestBookingRepository_FindByID(t *testing.T) {
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
				return id.Booking
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
				return id.Booking
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Booking().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Booking().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestBookingRepository_Delete(t *testing.T) {
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
				booking := model.TestBooking()
				booking.Employee.EmployeeID = id.Employee
				booking.Pet.PetID = id.Pet
				booking.Seat.SeatID = id.Seat
				id, _ := testStore.Booking().Create(booking)
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
				return 1
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Booking().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Booking().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestBookingRepository_Update(t *testing.T) {
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Booking
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Booking {
				testStore.Open()

				booking := model.TestBooking()
				booking.Employee.EmployeeID = id.Employee
				booking.Pet.PetID = id.Pet
				booking.Seat.SeatID = id.Seat
				booking.BookingID = id.Booking

				return booking
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.Booking {
				testStore.Open()

				booking := model.TestBooking()
				booking.Employee.EmployeeID = id.Employee
				booking.Pet.PetID = id.Pet
				booking.Seat.SeatID = id.Seat
				booking.BookingID = 0

				return booking
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Booking {
				testStore.Close()
				b := model.TestBooking()
				return b
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Booking().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Booking().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestBookingRepository_ModelFromDTO(t *testing.T) {
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.BookingDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.BookingDTO {
				testStore.Open()

				b := model.TestBookingDTO()
				b.PetID = id.Pet
				b.EmployeeID = id.Employee
				b.SeatID = id.Seat

				return b
			},
			isValid: true,
		},
		{
			name: "Invalid PetID",
			model: func() *model.BookingDTO {
				testStore.Open()

				b := model.TestBookingDTO()
				b.PetID = 0
				b.EmployeeID = id.Employee
				b.SeatID = id.Seat

				return b
			},
			isValid: false,
		},
		{
			name: "Invalid SeatID",
			model: func() *model.BookingDTO {
				testStore.Open()

				b := model.TestBookingDTO()
				b.PetID = id.Pet
				b.EmployeeID = id.Employee
				b.SeatID = 0

				return b
			},
			isValid: false,
		},
		{
			name: "Invalid EmployeeID",
			model: func() *model.BookingDTO {
				testStore.Open()

				b := model.TestBookingDTO()
				b.PetID = id.Pet
				b.EmployeeID = 0
				b.SeatID = id.Seat

				return b
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.BookingDTO {
				testStore.Close()
				b := model.TestBookingDTO()
				return b
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Booking().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Booking().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
