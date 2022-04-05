package store_test

import (
	"hotel/domain/model"
	"hotel/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeatRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Seat
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Seat {
				testStore.Open()

				seat := model.TestSeat()
				seat.Room.RoomID = id.Room

				return seat
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Seat {
				testStore.Close()
				return model.TestSeat()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Seat().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Seat().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSeatRepository_GetAll(t *testing.T) {
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
				result, err := tc.s().Seat().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Seat().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSeatRepository_FindByID(t *testing.T) {
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
				return id.Seat
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
				return id.Seat
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Seat().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Seat().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestSeatRepository_Delete(t *testing.T) {
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
				seat := model.TestSeat()
				seat.Room.RoomID = id.Room
				id, _ := testStore.Seat().Create(seat)
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
				return id.Seat
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Seat().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Seat().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestSeatRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Seat
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Seat {
				testStore.Open()

				seat := model.TestSeat()
				seat.SeatID = id.Seat
				seat.Room.RoomID = id.Room

				return seat
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.Seat {
				testStore.Open()

				seat := model.TestSeat()
				seat.SeatID = 0
				seat.Room.RoomID = id.Room

				return seat
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Seat {
				testStore.Close()

				seat := model.TestSeat()
				seat.SeatID = id.Seat
				seat.Room.RoomID = id.Room

				return seat
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Seat().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Seat().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestSeatRepository_ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.SeatDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.SeatDTO {
				testStore.Open()

				seat := model.TestSeatDTO()
				seat.RoomID = id.Room

				return seat
			},
			isValid: true,
		},
		{
			name: "Invalid RoomID",
			model: func() *model.SeatDTO {
				testStore.Open()

				seat := model.TestSeatDTO()
				seat.RoomID = 0

				return seat
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.SeatDTO {
				testStore.Close()
				seat := model.TestSeatDTO()
				return seat
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Seat().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Seat().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
