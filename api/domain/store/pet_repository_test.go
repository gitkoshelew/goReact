package store_test

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPetRepository_Create(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Pet
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Pet {
				testStore.Open()

				pet := model.TestPet()
				pet.Owner.UserID = id.User

				return pet
			},
			isValid: true,
		},
		{
			name: "DB closed",
			model: func() *model.Pet {
				testStore.Close()
				return model.TestPet()
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Pet().Create(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Pet().Create(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestPetRepository_GetAll(t *testing.T) {
	teardown()
	defer teardown()
	store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		s       func() *store.Store
		isValid bool
	}{
		{
			name: "valid",
			s: func() *store.Store {
				testStore.Open()
				return testStore
			},
			isValid: true,
		},
		{
			name: "DB closed",
			s: func() *store.Store {
				st := testStore
				st.Close()
				return st
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := tc.s().Pet().GetAll()
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := tc.s().Pet().GetAll()
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestPetRepository_FindByID(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()
				return id.Pet
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.Pet
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Pet().FindByID(tc.id())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Pet().FindByID(tc.id())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestPetRepository_Delete(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		id      func() int
		isValid bool
	}{
		{
			name: "valid",
			id: func() int {
				testStore.Open()
				pet := model.TestPet()
				pet.Owner.UserID = id.User
				id, _ := testStore.Pet().Create(pet)
				return *id
			},
			isValid: true,
		},
		{
			name: "Invalid ID",
			id: func() int {
				testStore.Open()
				return 0
			},
			isValid: false,
		},
		{
			name: "DB closed",
			id: func() int {
				testStore.Close()
				return id.Pet
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Pet().Delete(tc.id())
				assert.NoError(t, err)
				testStore.Close()
			} else {
				err := testStore.Pet().Delete(tc.id())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestPetRepository_Update(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.Pet
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.Pet {
				testStore.Open()

				pet := model.TestPet()
				pet.PetID = id.Pet
				pet.Owner.UserID = id.User

				return pet
			},
			isValid: true,
		},
		{
			name: "invalid ID",
			model: func() *model.Pet {
				testStore.Open()

				pet := model.TestPet()
				pet.PetID = 0
				pet.Owner.UserID = id.User

				return pet
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.Pet {
				testStore.Close()

				pet := model.TestPet()
				pet.PetID = id.Pet
				pet.Owner.UserID = id.User
				return pet
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				err := testStore.Pet().Update(tc.model())
				testStore.Close()
				assert.NoError(t, err)
			} else {
				err := testStore.Pet().Update(tc.model())
				testStore.Close()
				assert.Error(t, err)
			}
		})
	}
}

func TestPetRepository_ModelFromDTO(t *testing.T) {
	teardown()
	defer teardown()
	id := store.FillDB(t, testStore)

	testCases := []struct {
		name    string
		model   func() *model.PetDTO
		isValid bool
	}{
		{
			name: "valid",
			model: func() *model.PetDTO {
				testStore.Open()

				pet := model.TestPetDTO()
				pet.OwnerID = id.User

				return pet
			},
			isValid: true,
		},
		{
			name: "invalid OwnerId",
			model: func() *model.PetDTO {
				testStore.Open()

				pet := model.TestPetDTO()
				pet.OwnerID = 0

				return pet
			},
			isValid: false,
		},
		{
			name: "DB closed",
			model: func() *model.PetDTO {
				testStore.Close()
				p := model.TestPetDTO()
				return p
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				result, err := testStore.Pet().ModelFromDTO(tc.model())
				testStore.Close()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				result, err := testStore.Pet().ModelFromDTO(tc.model())
				testStore.Close()
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
