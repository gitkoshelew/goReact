package store

// Pet struct
type Pet struct {
	PetID       int     `json:"petId"`
	Name        string  `json:"name"`
	Type        PetType `json:"petType"`
	Weight      float32 `json:"weight"`
	Diesieses   string  `json:"diesieses"`
	Owner       Client
	PetPhotoURL string `json:"petPhotoUrl"`
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
