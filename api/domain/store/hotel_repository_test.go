package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		h, err := s.Hotel().Create(model.TestHotel())
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}

func TestHotelRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id delete", func(t *testing.T) {
		id := 0
		err := s.Hotel().Delete(id)
		assert.Error(t, err)
	})
	h := model.TestHotel()
	t.Run("valid id delete", func(t *testing.T) {
		_, err := s.Hotel().Create(h)
		err = s.Hotel().Delete(h.HotelID)
		assert.NoError(t, err)
	})
}

func TestHotelRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id FindByID", func(t *testing.T) {
		id := 2
		_, err := s.Hotel().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid id FindByID", func(t *testing.T) {
		h := model.TestHotel()
		_, err := s.Hotel().Create(h)
		h, err = s.Hotel().FindByID(h.HotelID)
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}

func TestHotelRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("get all valid", func(t *testing.T) {
		h, err := s.Hotel().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}

func TestHotelRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid update", func(t *testing.T) {
		h := model.TestHotel()
		_, err := s.Hotel().Create(h)

		h.Name = "Minsk"
		h.Address = "Minsk"

		err = s.Hotel().Update(h)
		assert.NoError(t, err)
	})
}
