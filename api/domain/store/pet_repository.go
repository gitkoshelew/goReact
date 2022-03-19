package store

import (
	"goReact/domain/model"
)

// PetRepository ...
type PetRepository struct {
	Store *Store
}

// Create pet and save it to DB
func (r *PetRepository) Create(p *model.PetDTO) (*model.Pet, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO pet (name, type, weight, diseases, user_id , photo) VALUES ($1, $2, $3, $4, $5 ,$6) RETURNING id",
		p.Name,
		p.Type,
		p.Weight,
		p.Diseases,
		p.OwnerID,
		p.PhotoURL,
	).Scan(&p.PetID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating pet. Err msg:%v.", err)
		return nil, err
	}

	pet, err := r.ModelFromDTO(p)
	if err != nil {
		return nil, err
	}

	r.Store.Logger.Infof("Pet with id %d was created.", p.PetID)

	return pet, nil
}

// GetAll returns all pets
func (r *PetRepository) GetAll() (*[]model.PetDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM pet")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all pets. Err msg: %v", err)
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
			r.Store.Logger.Errorf("Error occured while getting all pets. Err msg: %v", err)
			continue
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

// FindByID searchs and returns petDTO by ID
func (r *PetRepository) FindByID(id int) (*model.Pet, error) {
	petDTO := &model.PetDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE id = $1",
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

	pet, err := r.ModelFromDTO(petDTO)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

// Delete pet from DB by ID
func (r *PetRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM pet WHERE id = $1", id)
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
		r.Store.Logger.Errorf("Error occured while deleting pet. Err msg:%v.", err)
		return ErrNoRowsAffected
	}
	r.Store.Logger.Infof("Pet deleted, rows affectet: %d", result)
	return nil
}

// Update pet from DB
func (r *PetRepository) Update(p *model.Pet) error {
	result, err := r.Store.Db.Exec(
		"UPDATE pet SET name = $1, type = $2, weight = $3, diseases = $4, user_id = $5 , user_id = $6 WHERE id = $7",
		p.Name,
		string(p.Type),
		p.Weight,
		p.Diseases,
		p.Owner,
		p.PhotoURL,
		p.PetID,
	)
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
		r.Store.Logger.Errorf("Error occured while updating pet. Err msg:%v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Pet with id %d was updated", p.PetID)
	return nil
}

// ModelFromDTO ...
func (r *PetRepository) ModelFromDTO(dto *model.PetDTO) (*model.Pet, error) {
	user, err := r.Store.User().FindByID(dto.OwnerID)
	if err != nil {
		return nil, err
	}

	return &model.Pet{
		PetID:    dto.PetID,
		Name:     dto.Name,
		Type:     model.PetType(dto.Type),
		Weight:   dto.Weight,
		Diseases: dto.Diseases,
		Owner:    *user,
		PhotoURL: dto.PhotoURL,
	}, nil
}
