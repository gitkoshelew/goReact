package model

// Pet struct
type Pet struct {
	PetID       int     `json:"petId"`
	Name        string  `json:"name"`
	Type        PetType `json:"petType"`
	Weight      float32 `json:"weight"`
	Diesieses   string  `json:"diesieses"`
	Owner       User
	PetPhotoURL string `json:"petPhotoUrl"`
}

// PetType ...
type PetType string

// PetType constants
const (
	PetTypeCat PetType = "cat"
	PetTypeDog PetType = "dog"
)
