package entity

type PetType string

const (
	PetTypeCat PetType = "cat"
	PetTypeDog PetType = "dog"
)

func (pt PetType) IsValid() bool {
	switch pt {
	case PetTypeCat, PetTypeDog:
		return true
	}
	return false
}
