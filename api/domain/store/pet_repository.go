package store

import (
	"fmt"
	"goReact/domain/model"
)

// PetRepository ...
type PetRepository struct {
	Store *Store
}

// Create pet and save it to DB
func (r *PetRepository) Create(p *model.Pet) (*int, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO pet (name, type, weight, diseases, user_id , photoURL) VALUES ($1, $2, $3, $4, $5 ,$6) RETURNING pet_id",
		p.Name,
		p.Type,
		p.Weight,
		p.Diseases,
		p.Owner.UserID,
		p.PhotoURL,
	).Scan(&p.PetID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating pet. Err msg:%v.", err)
		return nil, err
	}

	r.Store.Logger.Infof("Pet with id %d was created.", p.PetID)

	return &p.PetID, nil
}

// GetAll returns all pets
func (r *PetRepository) GetAll() (*[]model.PetDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM pet")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all pets. Err msg: %v", err)
		return nil, err
	}
	pets := []model.PetDTO{}

	for rows.Next() {
		pet := model.PetDTO{}
		err := rows.Scan(
			&pet.PetID,
			&pet.Name,
			&pet.Type,
			&pet.Weight,
			&pet.Diseases,
			&pet.OwnerID,
			&pet.PhotoURL,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while getting all pets. Err msg: %v", err)
			continue
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

// FindByID searchs and returns petDTO by ID
func (r *PetRepository) FindByID(id int) (*model.PetDTO, error) {
	petDTO := &model.PetDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE pet_id = $1",
		id).Scan(
		&petDTO.PetID,
		&petDTO.Name,
		&petDTO.Type,
		&petDTO.Weight,
		&petDTO.Diseases,
		&petDTO.OwnerID,
		&petDTO.PhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting pet by id. Err msg:%v.", err)
		return nil, err
	}

	return petDTO, nil
}

// Delete pet from DB by ID
func (r *PetRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM pet WHERE pet_id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting pet. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting pet. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting pet. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}
	r.Store.Logger.Infof("Pet deleted, rows affectet: %d", result)
	return nil
}

// Update pet from DB
func (r *PetRepository) Update(p *model.Pet) error {
	name := "name"
	if p.Name != "" {
		name = fmt.Sprintf("'%s'", p.Name)
	}
	petType := "type"
	if p.Type != "" {
		petType = fmt.Sprintf("'%s'", string(p.Type))
	}
	weight := "weight"
	if p.Weight != 0 {
		weight = fmt.Sprintf("%v", p.Weight)
	}
	diseases := "diseases"
	if p.Diseases != "" {
		diseases = fmt.Sprintf("'%s'", p.Diseases)
	}
	ownerID := "user_id"
	if p.Owner.UserID != 0 {
		ownerID = fmt.Sprintf("%d", p.Owner.UserID)
	}
	photoURL := "photoURL"
	if p.PhotoURL != "" {
		photoURL = fmt.Sprintf("'%s'", p.PhotoURL)
	}

	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE pet SET 
		name = %s, type = %s, weight = %s, diseases = %s, user_id = %s , photoURL = %s 
		WHERE pet_id = $1`,
		name,
		petType,
		weight,
		diseases,
		ownerID,
		photoURL,
	), p.PetID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating pet. Err msg:%v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating pet. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating pet. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Pet with id %d was updated", p.PetID)
	return nil
}

// ModelFromDTO ...
func (r *PetRepository) ModelFromDTO(dto *model.PetDTO) (*model.Pet, error) {
	userDTO, err := r.Store.User().FindByID(dto.OwnerID)
	if err != nil {
		return nil, err
	}
	u, err := r.Store.User().ModelFromDTO(userDTO)
	if err != nil {
		return nil, err
	}

	return &model.Pet{
		PetID:    dto.PetID,
		Name:     dto.Name,
		Type:     model.PetType(dto.Type),
		Weight:   dto.Weight,
		Diseases: dto.Diseases,
		Owner:    *u,
		PhotoURL: dto.PhotoURL,
	}, nil
}
