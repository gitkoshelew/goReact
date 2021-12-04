package entity

import "goReact/domain/dto"

// Pet struct
type Pet struct {
	PetID     int     `json:"petId"`
	Name      string  `json:"name"`
	Type      PetType `json:"petType"`
	OwnerID   int     `json:"ownerId"`
	Weight    float32 `json:"weight"`
	Diesieses string  `json:"diesieses"`
}

// SetName sets Pets Name
func (p *Pet) SetName(s string) {
	p.Name = s
}

// SetType sets Pets Type
func (p *Pet) SetType(pt PetType) {
	p.Type = pt
}

// SetWeight sets Pets Weight
func (p *Pet) SetWeight(f float32) {
	p.Weight = f
}

// SetDiesieses sets Pets Diesieses
func (p *Pet) SetDiesieses(s string) {
	p.Diesieses = s
}

// SetOwnerID sets Pets Owner
func (p *Pet) SetOwnerID(i int) {
	p.OwnerID = i
}

// GetPetByID returns Pet by id from storage
func GetPetByID(id int) Pet {
	var pet Pet
	for _, p := range GetPets() {
		if id == p.PetID {
			pet = p
		}
	}
	return pet
}

// GetPetsByID returns []Pets by []ids from storage
func GetPetsByID(ids []int) []Pet {
	var pets []Pet
	for _, id := range ids {
		pets = append(pets, GetPetByID(id))
	}
	return pets
}

// PetToDto makes DTO from Pet object
func PetToDto(p Pet) dto.Pet {
	return dto.Pet{
		PetID:     p.PetID,
		Name:      p.Name,
		Type:      string(p.Type),
		OwnerID:   p.OwnerID,
		Weight:    p.Weight,
		Diesieses: p.Diesieses,
	}
}
