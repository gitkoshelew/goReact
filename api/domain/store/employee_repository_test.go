package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("employee")
	t.Run("valid", func(t *testing.T) {
		e, err := s.Employee().Create(model.TestEmployee(t))
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("employee")

	t.Run("invalid id", func(t *testing.T) {
		id := 2
		err := s.Employee().Delete(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee(t)
		_, err := s.Employee().Create(e)
		err = s.Employee().Delete(e.EmployeeID)
		assert.NoError(t, err)
	})
}

func TestEmployeeRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("employee")
	t.Run("invalid id", func(t *testing.T) {
		id := 2
		_, err := s.Employee().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		e := model.TestEmployee(t)
		e, err := s.Employee().FindByID(e.EmployeeID)
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("employee")
	t.Run("valid ", func(t *testing.T) {
		e, err := s.Employee().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}

