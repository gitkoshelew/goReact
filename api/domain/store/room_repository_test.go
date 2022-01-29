package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		r := model.TestRoom()
		h, err := s.Hotel().Create(model.TestHotel())
		r.Hotel = *h
		r, err = s.Room().Create(r)
		assert.NoError(t, err)
		assert.NotNil(t, r)
	})

}

func TestRoomRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("invalid delete id", func(t *testing.T) {
		id := 2

		err := s.Room().Delete(id)
		assert.Error(t, err)
	})
	t.Run("valid delete id", func(t *testing.T) {
		r := model.TestRoom()
		h, err := s.Hotel().Create(model.TestHotel())
		r.Hotel = *h
		_, err = s.Room().Create(r)
		err = s.Room().Delete(r.RoomID)
		assert.NoError(t, err)
	})
}

func TestRoomRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("invalid find id", func(t *testing.T) {
		id := -1

		_, err := s.Room().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid find id", func(t *testing.T) {
		r := model.TestRoom()
		h, err := s.Hotel().Create(model.TestHotel())
		r.Hotel = *h
		_, err = s.Room().Create(r)
		rDTO := model.RoomDTO{
			RoomID:       r.RoomID,
			RoomNumber:   r.RoomNumber,
			PetType:      r.PetType,
			HotelID:      r.Hotel.HotelID,
			RoomPhotoURL: r.RoomPhotoURL,
		}

		room, err := s.Room().FindByID(rDTO.RoomID)
		assert.NoError(t, err)
		assert.NotNil(t, room)
	})
}

func TestRoomRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid get all", func(t *testing.T) {
		r, err := s.Room().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, r)
	})
}

func TestRoomRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid update", func(t *testing.T) {
		r := model.TestRoom()
		h, err := s.Hotel().Create(model.TestHotel())
		r.Hotel = *h
		r, err = s.Room().Create(r)

		r.RoomNumber = 2222
		r.PetType = "dog"
		r.RoomPhotoURL = "//photo//2"

		err = s.Room().Update(r)
		assert.NoError(t, err)
	})
}

func TestRoomRepository_GetAllPagination(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid get all panigation", func(t *testing.T) {
		p := model.TestPage()
		r, err := s.Room().GetAllPagination(p)
		assert.NoError(t, err)
		assert.NotNil(t, r)
	})
	t.Run("invalid get all panigation", func(t *testing.T) {
		p := model.TestPage()
		p.PageNumber = -1
		p.PageSize = -10
		r, err := s.Room().GetAllPagination(p)
		assert.Error(t, err)
		assert.Nil(t, r)
	})
}
