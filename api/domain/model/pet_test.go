package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPet_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.PetDTO
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.PetDTO {
				return model.TestPetDTO()
			},
			isValid: true,
		},
		{
			name: "Invalid Name",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Name = "Name@123"
				return p
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Name = ""
				return p
			},
			isValid: false,
		},
		{
			name: "Invalid Type",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Type = "Invalid"
				return p
			},
			isValid: false,
		},
		{
			name: "Big Weight",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Weight = 999
				return p
			},
			isValid: false,
		},
		{
			name: "Small Weight",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Weight = 0.001
				return p
			},
			isValid: false,
		},
		{
			name: "SQL Diseases",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.Diseases = "AL (&-Te R"
				return p
			},
			isValid: false,
		},
		{
			name: "Empty OwnerID",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.OwnerID = 0
				return p
			},
			isValid: false,
		},
		{
			name: "invalid OwnerID",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.OwnerID = -1
				return p
			},
			isValid: false,
		},
		{
			name: "SQL PhotoURL",
			u: func() *model.PetDTO {
				p := model.TestPetDTO()
				p.PhotoURL = "AL (&-Te R"
				return p
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
