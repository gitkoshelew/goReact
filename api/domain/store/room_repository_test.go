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

	r, err := s.Room().Create(model.TestRoom())
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestRoomRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	id := 2
	t.Run("invalid ID", func(t *testing.T) {
		err := s.Room().Delete(id)
		assert.Error(t, err)
	})

	r := model.TestRoom()

	t.Run("valid ID", func(t *testing.T) {
		_, err := s.Room().Create(r)
		err = s.Room().Delete(r.RoomID)
		assert.NoError(t, err)
	})

}

func TestRoomRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	id := 2
	t.Run("invalid ID", func(t *testing.T) {
		_, err := s.Room().FindByID(id)
		assert.Error(t, err)
	})

	r := model.TestRoom()

	t.Run("valid ID", func(t *testing.T) {
		r, err := s.Room().FindByID(r.RoomID)
		assert.NoError(t, err)
		assert.NotNil(t, r)
	})
}

func TestRoomRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	r, err := s.Room().GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestRoomRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	r := model.Room{
		RoomID: 2,
		RoomNumber: 2,
		PetType: model.PetTypeDog,
		Hotel: *model.TestHotel(),
		RoomPhotoURL: "//photo//2",
	}

	err := s.Room().Update(&r)
	assert.Error(t, err)
	r1 := model.TestRoom()
	err = s.Room().Update(r1)

	assert.NoError(t, err)
	assert.NotNil(t, r)
}
