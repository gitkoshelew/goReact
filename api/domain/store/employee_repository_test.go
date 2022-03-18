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
		u := model.TestUser()
		h := model.TestHotel()
		e.User = *u
		e.Hotel = *h
		id, err := s.Employee().Create(&model.Employee{
			EmployeeID: e.EmployeeID,
			User:       *u,
			Hotel:      *h,
			Position:   e.Position,
		})
		assert.NoError(t, err)
		assert.NotNil(t, id)
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
		u := model.TestUser()
		h := model.TestHotel()
		e.User = *u
		e.Hotel = *h
		id, _ := s.Employee().Create(&model.Employee{
			EmployeeID: e.EmployeeID,
			User:       *u,
			Hotel:      *h,
			Position:   e.Position,
		})
		err := s.Employee().Delete(*id)
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
		u := model.TestUser()
		h := model.TestHotel()
		e.User = *u
		e.Hotel = *h
		id, _ := s.Employee().Create(&model.Employee{
			EmployeeID: e.EmployeeID,
			User:       *u,
			Hotel:      *h,
			Position:   e.Position,
		})
		eDTO, err := s.Employee().FindByID(*id)
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
		u := model.TestUser()
		h := model.TestHotel()
		e.User = *u
		e.Hotel = *h
		id, _ := s.Employee().Create(&model.Employee{
			EmployeeID: e.EmployeeID,
			User:       *u,
			Hotel:      *h,
			Position:   e.Position,
		})
		eDTO, err := s.Employee().FindByID(*id)
		assert.NoError(t, err)
		assert.NotNil(t, eDTO)
	})
}
