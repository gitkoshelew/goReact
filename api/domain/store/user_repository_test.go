package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	u, err := s.User().Create(model.TestUser())
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	id := -1

	err := s.User().Delete(id)
	assert.Error(t, err)

	u := model.TestUser()

	_, err = s.User().Create(u)

	err = s.User().Delete(u.UserID)
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("Invalid email", func(t *testing.T) {
		u, err := s.User().FindByEmail("invalid@example.org")
		assert.Error(t, err)
		assert.Nil(t, u)
	})

	u := model.TestUser()
	u, _ = s.User().Create(u)

	t.Run("valid email", func(t *testing.T) {
		u, err := s.User().FindByEmail(u.Email)
		assert.NoError(t, err)
		assert.NotNil(t, u)
	})

}

func TestUserRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
            
	t.Run("Invalid ID", func(t *testing.T) {
		u, err := s.User().FindByID(-1)
		assert.Error(t, err)
		assert.Nil(t, u)

	})

	u := model.TestUser()
	u, _ = s.User().Create(u)

	t.Run("Valid ID", func(t *testing.T) {
		u, err := s.User().FindByID(u.UserID)

		assert.NoError(t, err)
		assert.NotNil(t, u)

	})

}

func TestUserRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	u, err := s.User().GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_VerifyEmail(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("Invalid user ID", func(t *testing.T) {
		err := s.User().VerifyEmail(-1)

		assert.Error(t, err)
	})

	u := model.TestUser()
	user, _ := s.User().Create(u)

	t.Run("Valid", func(t *testing.T) {
		err := s.User().VerifyEmail(user.UserID)
		assert.NoError(t, err)
	})

}
