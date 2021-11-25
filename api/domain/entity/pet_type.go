package entity

// PetType ...
type PetType string

// PetType constants
const (
	PetTypeCat PetType = "cat"
	PetTypeDog PetType = "dog"
)

//IsValid checks is Pet Type is valid
func (p PetType) IsValid() bool {
	switch p {
	case PetTypeCat, PetTypeDog:
		return true
	}
	return false
}
