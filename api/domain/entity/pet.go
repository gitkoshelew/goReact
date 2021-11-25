package entity

// Pet struct
type Pet struct {
	PetID     int
	Name      string
	Type      PetType
	OwnerID   int
	Weight    float32
	Diesieses string
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
