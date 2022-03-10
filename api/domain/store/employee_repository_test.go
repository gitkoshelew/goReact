package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee()
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u.ModelFromDTO()
		e.Hotel = *h
		e, err := s.Employee().Create(&model.EmployeeDTO{
			EmployeeID: e.EmployeeID,
			UserID:     e.User.UserID,
			HotelID:    e.Hotel.HotelID,
			Position:   string(e.Position),
		})
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
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u.ModelFromDTO()
		e.Hotel = *h
		e, _ = s.Employee().Create(&model.EmployeeDTO{
			EmployeeID: e.EmployeeID,
			UserID:     e.User.UserID,
			HotelID:    e.Hotel.HotelID,
			Position:   string(e.Position),
		})
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
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u.ModelFromDTO()
		e.Hotel = *h
		e, _ = s.Employee().Create(&model.EmployeeDTO{
			EmployeeID: e.EmployeeID,
			UserID:     e.User.UserID,
			HotelID:    e.Hotel.HotelID,
			Position:   string(e.Position),
		})
		eDTO, err := s.Employee().FindByID(e.EmployeeID)
		e, err = s.EmployeeRepository.ModelFromDTO(eDTO)
		assert.NoError(t, err)
		assert.NotNil(t, e)
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

func TestEmployeeRepository_FindByUserID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id", func(t *testing.T) {
		id := -1
		_, err := s.Employee().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee()
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u.ModelFromDTO()
		e.Hotel = *h
		e, _ = s.Employee().Create(&model.EmployeeDTO{
			EmployeeID: e.EmployeeID,
			UserID:     e.User.UserID,
			HotelID:    e.Hotel.HotelID,
			Position:   string(e.Position),
		})
		eDTO, err := s.Employee().FindByID(e.EmployeeID)
		e, err = s.EmployeeRepository.ModelFromDTO(eDTO)
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}
