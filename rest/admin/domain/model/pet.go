package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Pet struct
type Pet struct {
	PetID    int     `json:"petId"`
	Name     string  `json:"name"`
	Type     PetType `json:"petType"`
	Weight   float32 `json:"weight"`
	Diseases string  `json:"diseases,omitempty"`
	Owner    User
	PhotoURL string `json:"photoUrl,omitempty"`
}

// PetDTO struct
type PetDTO struct {
	PetID    int     `json:"petId"`
	Name     string  `json:"name"`
	Type     string  `json:"petType"`
	Weight   float32 `json:"weight"`
	Diseases string  `json:"diseases,omitempty"`
	OwnerID  int     `json:"userId"`
	PhotoURL string  `json:"photoUrl,omitempty"`
}

// PetType ...
type PetType string

// PetType constants
const (
	PetTypeCat PetType = "cat"
	PetTypeDog PetType = "dog"
)

// Validate ...
func (p *PetDTO) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.By(IsSQL), validation.Length(2, 20)),
		validation.Field(&p.Type, validation.Required, validation.By(IsPetType)),
		validation.Field(&p.Weight, validation.Required, validation.Min(0.01), validation.Max(49.9)),
		validation.Field(&p.Diseases, validation.By(IsSQL)),
		validation.Field(&p.OwnerID, validation.Required, validation.By(IsValidID)),
		validation.Field(&p.PhotoURL, validation.By(IsSQL)),
	)
}
