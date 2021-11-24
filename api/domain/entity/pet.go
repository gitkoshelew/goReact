package entity

import "fmt"

// Pet ...
type Pet struct {
	PetID     int
	Name      string
	Type      string
	OwnerID   int
	Weight    float32
	Diesieses string
}

func (p *Pet) getInfo() string {
	return fmt.Sprintf("Pet ID: %d\n"+
		"Name: %s\n"+
		"Type: %s\n"+
		"OwnerID: %d\n"+
		"Weight: %f\n"+
		"Diesieses: %s\n",
		p.PetID, p.Name, p.Type, p.OwnerID, p.Weight, p.Diesieses)
}

func (p *Pet) setName(s string) {
	p.Name = s
}

func (p *Pet) setType(s string) {
	p.Type = s
}

func (p *Pet) setWeight(f float32) {
	p.Weight = f
}

func (p *Pet) setDiesieses(s string) {
	p.Diesieses = s
}

func (p *Pet) setOwnerID(i int) {
	p.OwnerID = i
}
