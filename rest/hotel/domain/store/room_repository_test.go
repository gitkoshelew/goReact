package store_test

import (
	"hotel/domain/model"
	"hotel/domain/store"
	"hotel/pkg/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Room
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Room {
				testStore.Open()

				room := model.TestRoom()
				room.Hotel.HotelID = id.Hotel

				return room
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Room {
				testStore.Close()
				return model.TestRoom()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Room().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Room().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestRoomRepository_GetAll(t *testing.T) {
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
				result, err := tc.s().Room().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Room().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestRoomRepository_FindByID(t *testing.T) {
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
				return id.Room
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
				return id.Room
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Room().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Room().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestRoomRepository_Delete(t *testing.T) {
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
				room := model.TestRoom()
				room.Hotel.HotelID = id.Hotel
				id, _ := testStore.Room().Create(room)
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
				return id.Room
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Room().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Room().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestRoomRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Room
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Room {
				testStore.Open()

				room := model.TestRoom()
				room.RoomID = id.Room
				room.Hotel.HotelID = id.Hotel

				return room
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.Room {
				testStore.Open()

				room := model.TestRoom()
				room.RoomID = 0
				room.Hotel.HotelID = id.Hotel

				return room
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Room {
				testStore.Close()

				room := model.TestRoom()
				room.RoomID = id.Room
				room.Hotel.HotelID = id.Hotel

				return room
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Room().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Room().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestRoomRepository_ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.RoomDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.RoomDTO {
				testStore.Open()

				room := model.TestRoomDTO()
				room.HotelID = id.Hotel

				return room
			},
			isValid: true,
		},
		{
			name: "Invalid HotelID",
			model: func() *model.RoomDTO {
				testStore.Open()

				room := model.TestRoomDTO()
				room.HotelID = 0

				return room
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.RoomDTO {
				testStore.Close()
				b := model.TestRoomDTO()
				return b
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Room().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Room().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestRoomRepository_GetAllPagination(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *pagination.Page
		isValid bool
	}{
		{
			name: "valid",
			model: func() *pagination.Page {
				testStore.Open()

				page := model.TestPage()

				return page
			},
			isValid: true,
		},
		{
			name: "invalid",
			model: func() *pagination.Page {
				testStore.Open()

				page := model.TestPage()
				page.PageNumber = -1
				page.PageSize = -1

				return page
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *pagination.Page {
				testStore.Close()

				page := model.TestPage()

				return page
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Room().GetAllPagination(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Room().GetAllPagination(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestRoomRepository_GetTotalRows(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		store   func() *store.Store
		isValid bool
	}{
		{
			name: "valid",
			store: func() *store.Store {
				testStore.Open()
				return testStore
			},
			isValid: true,
		},
		{
			name: "DB closed",
			store: func() *store.Store {
				testStore.Close()

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
				_, err := tc.store().Room().GetTotalRows()
				testStore.Close()
				assert.NoError(t, err)
			} else {
				_, err := tc.store().Room().GetTotalRows()
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}
