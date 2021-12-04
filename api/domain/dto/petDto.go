package dto

// Pet DTO
type Pet struct {
	PetID     int     `json:"petId"`
	Name      string  `json:"name"`
	Type      string  `json:"petType"`
	OwnerID   int     `json:"ownerId"`
	Weight    float32 `json:"weight"`
	Diesieses string  `json:"diesieses"`
}
