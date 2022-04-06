package store_test

import (
	"admin/domain/model"
	"admin/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository_Create(t *testing.T) {
	teardown()
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
				return id.Booking
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
	teardown()
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

				booking := model.TestBooking()
				booking.Employee.EmployeeID = id.Employee
				booking.Pet.PetID = id.Pet
				booking.Seat.SeatID = id.Seat
				booking.BookingID = id.Booking

				return booking
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
	teardown()
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

				booking := model.TestBookingDTO()
				booking.PetID = id.Pet
				booking.EmployeeID = id.Employee
				booking.SeatID = id.Seat

				return booking
			},
			isValid: true,
		},
		{
			name: "Invalid PetID",
			model: func() *model.BookingDTO {
				testStore.Open()

				booking := model.TestBookingDTO()
				booking.PetID = 0
				booking.EmployeeID = id.Employee
				booking.SeatID = id.Seat

				return booking
			},
			isValid: false,
		},
		{
			name: "Invalid SeatID",
			model: func() *model.BookingDTO {
				testStore.Open()

				booking := model.TestBookingDTO()
				booking.PetID = id.Pet
				booking.EmployeeID = id.Employee
				booking.SeatID = 0

				return booking
			},
			isValid: false,
		},
		{
			name: "Invalid EmployeeID",
			model: func() *model.BookingDTO {
				testStore.Open()

				booking := model.TestBookingDTO()
				booking.PetID = id.Pet
				booking.EmployeeID = 0
				booking.SeatID = id.Seat

				return booking
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
