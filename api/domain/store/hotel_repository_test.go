package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Hotel
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Hotel {
				testStore.Open()

				hotel := model.TestHotel()

				return hotel
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Hotel {
				testStore.Close()
				return model.TestHotel()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Hotel().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Hotel().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestHotelRepository_GetAll(t *testing.T) {
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
				result, err := tc.s().Hotel().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Hotel().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestHotelRepository_FindByID(t *testing.T) {
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
				return id.Hotel
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
				return id.Hotel
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Hotel().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Hotel().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestHotelRepository_Delete(t *testing.T) {
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
				hotel := model.TestHotel()
				id, _ := testStore.Hotel().Create(hotel)
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
				return id.Hotel
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Hotel().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Hotel().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestHotelRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Hotel
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Hotel {
				testStore.Open()

				hotel := model.TestHotel()
				hotel.HotelID = id.Hotel

				return hotel
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.Hotel {
				testStore.Open()

				hotel := model.TestHotel()

				return hotel
			},
			isValid: false,
		},
		{
			name: "nil Coordinates",
			model: func() *model.Hotel {
				testStore.Open()

				hotel := model.TestHotel()
				hotel.HotelID = id.Hotel
				hotel.Coordinates = nil

				return hotel
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Hotel {
				testStore.Close()

				hotel := model.TestHotel()

				return hotel
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Hotel().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Hotel().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}
