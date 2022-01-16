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

	u, err := s.User().Create(model.TestUser())
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	id := 1

	err := s.User().Delete(id)
	assert.Error(t, err)

	u := model.TestUser()
	u.UserID = id

	_, err = s.User().Create(u)

	err = s.User().Delete(id)
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")
	email := "email@example.org"

	t.Run("FindByEmail: Invalid email", func(t *testing.T) {
		_, err := s.User().FindByEmail(email)
		assert.Error(t, err)
	})

	u := model.TestUser()
	u.Email = email

	t.Run("FindByEmail: valid email", func(t *testing.T) {
		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)
	})

}

func TestUserRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")
	id := 1

	_, err := s.User().FindByID(id)
	assert.Error(t, err)

	u := model.TestUser()
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

func TestUserRepository_VerifyEmail(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	defer teardown("users")

	u := model.TestUser()
	user, _ := s.User().Create(u)

	t.Run("Valid", func(t *testing.T) {
		err := s.User().VerifyEmail(user.UserID)

		assert.NoError(t, err)
	})

	t.Run("Invalid user ID", func(t *testing.T) {
		err := s.User().VerifyEmail(-1)

		assert.Error(t, err)
	})

}
