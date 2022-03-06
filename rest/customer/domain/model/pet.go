package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Pet struct
type Pet struct {
	PetID       int     `json:"petId"`
	Name        string  `json:"name"`
	Type        PetType `json:"petType"`
	Weight      float32 `json:"weight"`
	Diseases    string  `json:"diseases"`
	Owner       User
	PetPhotoURL string `json:"petPhotoUrl"`
}

// PetDTO struct
type PetDTO struct {
	PetID       int     `json:"petId"`
	Name        string  `json:"name"`
	Type        PetType `json:"petType"`
	Weight      float32 `json:"weight"`
	Diseases    string  `json:"diseases"`
	OwnerID     int     `json:"userId"`
	PetPhotoURL string  `json:"petPhotoUrl"`
}

// PetType ...
type PetType string

// PetType constants
const (
	PetTypeCat PetType = "cat"
	PetTypeDog PetType = "dog"
)

// Validate ...
func (p *Pet) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Type, validation.Required, validation.By(IsPetType)),
		validation.Field(&p.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(1, 20)),
		validation.Field(&p.Weight, validation.Required, validation.Min(0.1), validation.Max(49.9)),
		validation.Field(&p.Owner, validation.Required),
	)
}
