package model_test

import (
	"goReact/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPet_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.Pet
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.Pet {
				return model.TestPet()
			},
			isValid: true,
		},
		{
			name: "Invalid Name",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Name = "Name@123"
				return p
			},
			isValid: false,
		},
		{
			name: "Empty Name",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Name = ""
				return p
			},
			isValid: false,
		},
		{
			name: "Invalid Type",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Type = "Invalid"
				return p
			},
			isValid: false,
		},
		{
			name: "Valid Type",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Type = "cat"
				return p
			},
			isValid: true,
		},
		{
			name: "Invalid Weight",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Weight = 999
				return p
			},
			isValid: false,
		},		
		{
			name: "Invalid Weight",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Weight = 0
				return p
			},
			isValid: false,
		},				
		{
			name: "Valid Weight",
			u: func() *model.Pet {
				p := model.TestPet()
				p.Weight = 2
				return p
			},
			isValid: true,
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

func TestPet_NewPet(t *testing.T) {
	p := model.TestPet()
	p.PetID = 1
	assert.NotNil(t, p)
}