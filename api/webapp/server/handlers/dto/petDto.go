package dto

// PetDto ...
type PetDto struct {
	PetID     int     `json:"petId"`
	Name      string  `json:"name"`
	Type      string  `json:"petType"`
	Weight    float32 `json:"weight"`
	Diesieses string  `json:"diesieses"`
	OwnerID   int     `json:"ownerId"`
}
