package entity

type PetId int

type Pet struct {
	Id   PetId
	Name string
	Type PetType
}
