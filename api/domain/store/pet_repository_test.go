package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPetRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		p, err := s.Pet().Create(model.TestPet())
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestPetRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id", func(t *testing.T) {
		id := 2
		err := s.Pet().Delete(id)
		assert.Error(t, err)
	})

	p := model.TestPet()
	t.Run("valid id", func(t *testing.T) {		
		_, err := s.Pet().Create(p)
		err = s.Pet().Delete(p.PetID)
		assert.NoError(t, err)
	})
}

func TestPetRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid Find by id", func(t *testing.T) {
		id := 2
		_, err := s.Pet().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid Find by id", func(t *testing.T) {
		p := model.TestPet()
		p, err := s.Pet().FindByID(p.PetID)
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestPetRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("Get all valid", func(t *testing.T) {
		p, err := s.Pet().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}
