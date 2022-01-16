package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("hotels")
	t.Run("valid", func(t *testing.T) {
		h, err := s.Hotel().Create(model.TestHotel(t))
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}

func TestHotelRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("hotels")
	t.Run("invalid id", func(t *testing.T) {
		id := 2
		err := s.Hotel().Delete(id)
		assert.Error(t, err)
	})
	t.Run("valid id", func(t *testing.T) {
		h := model.TestHotel(t)
		_, err := s.Hotel().Create(h)
		err = s.Hotel().Delete(h.HotelID)
		assert.NoError(t, err)
	})
}

func TestHotelRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("hotels")
	t.Run("invalid id", func(t *testing.T) {
		id := 2
		_, err := s.Hotel().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid id", func(t *testing.T) {
		h := model.TestHotel(t)
		h, err := s.Hotel().FindByID(h.HotelID)
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}

func TestHotelRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("hotels")
	t.Run("invalid id", func(t *testing.T) {
		h, err := s.Hotel().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, h)
	})
}
