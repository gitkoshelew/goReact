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
		u, err := s.User().Create(model.TestUser())
		p := model.TestPet()
		p.Owner = *u
		p, err = s.Pet().Create(&model.PetDTO{
			PetID:     p.PetID,
			Name:      p.Name,
			Type:      string(p.Type),
			Weight:    p.Weight,
			Diesieses: p.Diesieses,
			OwnerID:   p.Owner.UserID,
			PhotoURL:  p.PhotoURL,
		})
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestPetRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid Delete id", func(t *testing.T) {
		id := 2
		err := s.Pet().Delete(id)
		assert.Error(t, err)
	})

	t.Run("valid Delete id", func(t *testing.T) {
		u, err := s.User().Create(model.TestUser())
		p := model.TestPet()
		p.Owner = *u
		p, err = s.Pet().Create(&model.PetDTO{
			PetID:     p.PetID,
			Name:      p.Name,
			Type:      string(p.Type),
			Weight:    p.Weight,
			Diesieses: p.Diesieses,
			OwnerID:   p.Owner.UserID,
			PhotoURL:  p.PhotoURL,
		})
		err = s.Pet().Delete(p.PetID)
		assert.NoError(t, err)
	})
}

func TestPetRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid find by id", func(t *testing.T) {
		id := 2
		_, err := s.Pet().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid find by id", func(t *testing.T) {
		u, err := s.User().Create(model.TestUser())

		p := model.TestPet()
		p.Owner = *u
		p, err = s.Pet().Create(&model.PetDTO{
			PetID:     p.PetID,
			Name:      p.Name,
			Type:      string(p.Type),
			Weight:    p.Weight,
			Diesieses: p.Diesieses,
			OwnerID:   p.Owner.UserID,
			PhotoURL:  p.PhotoURL,
		})
		p, err = s.Pet().FindByID(p.PetID)
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

func TestPetRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid update", func(t *testing.T) {
		u, err := s.User().Create(model.TestUser())
		p := model.TestPet()
		p.Owner = *u
		p, err = s.Pet().Create(&model.PetDTO{
			PetID:     p.PetID,
			Name:      p.Name,
			Type:      string(p.Type),
			Weight:    p.Weight,
			Diesieses: p.Diesieses,
			OwnerID:   p.Owner.UserID,
			PhotoURL:  p.PhotoURL,
		})

		p.Name = "Sharik"
		p.Type = "dog"
		p.Weight = 2
		p.Diesieses = "Izjoga"
		p.PhotoURL = "/1/2/jpg"

		err = s.Pet().Update(&model.PetDTO{
			PetID:     p.PetID,
			Name:      p.Name,
			Type:      string(p.Type),
			Weight:    p.Weight,
			Diesieses: p.Diesieses,
			OwnerID:   p.Owner.UserID,
			PhotoURL:  p.PhotoURL,
		})
		assert.NoError(t, err)
	})
}
