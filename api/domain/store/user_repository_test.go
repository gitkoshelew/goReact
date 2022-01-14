package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	id := 1

	err := s.User().Delete(id)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.UserID = id

	_, err = s.User().Create(u)

	err = s.User().Delete(id)
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	email := "email@example.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")
	id := 1

	_, err := s.User().FindByID(id)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.UserID = id

	u, err = s.User().FindByID(id)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	u, err := s.User().GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
