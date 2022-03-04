package store_test

import (
	"hotel/domain/model"
	"hotel/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee()
		h, _ := s.Hotel().Create(model.TestHotel())

		e.Hotel = *h
		e, err := s.Employee().Create(e)
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("invalid id", func(t *testing.T) {
		id := -1
		err := s.Employee().Delete(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee()
		h, _ := s.Hotel().Create(model.TestHotel())

		e.Hotel = *h
		e, _ = s.Employee().Create(e)
		err := s.Employee().Delete(e.EmployeeID)
		assert.NoError(t, err)
	})
}

func TestEmployeeRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id", func(t *testing.T) {
		id := -1
		_, err := s.Employee().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee()

		h, _ := s.Hotel().Create(model.TestHotel())

		e.Hotel = *h
		e, _ = s.Employee().Create(e)
		eDTO, err := s.Employee().FindByID(e.EmployeeID)
		assert.NoError(t, err)
		assert.NotNil(t, eDTO)
	})
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid ", func(t *testing.T) {
		e, err := s.Employee().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}
