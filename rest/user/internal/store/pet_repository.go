package store

import (
	"errors"
	"fmt"
	"user/domain/model"
)

// PetRepository ...
type PetRepository struct {
	Store *Store
}

// Create pet and save it to DB
func (r *PetRepository) Create(p *model.Pet) (*model.Pet, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO pet (name, type, weight, diseases, user_id , photo) VALUES ($1, $2, $3, $4, $5 ,$6) RETURNING id",
		p.Name,
		string(p.Type),
		p.Weight,
		p.Diseases,
		p.Owner.UserID,
		p.PetPhotoURL,
	).Scan(&p.PetID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating pet. Err msg:%v.", err)
		return nil, err
	}
	r.Store.Logger.Info("Created pet with id = %d", p.PetID)

	return p, nil
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
			&pet.PetPhotoURL,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all pets. Err msg: %v", err)
			continue
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

// FindDTOByID searchs and returns petDTO by ID
func (r *PetRepository) FindDTOByID(id int) (*model.PetDTO, error) {
	pet := &model.PetDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE id = $1",
		id).Scan(
		&pet.PetID,
		&pet.Name,
		&pet.Type,
		&pet.Weight,
		&pet.Diseases,
		&pet.OwnerID,
		&pet.PetPhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting pet by id. Err msg:%v.", err)
		return nil, err
	}
	return pet, nil
}

// FindByID searchs and returns pet by ID
func (r *PetRepository) FindByID(id int) (*model.Pet, error) {
	pet := &model.Pet{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE id = $1",
		id).Scan(
		&pet.PetID,
		&pet.Name,
		&pet.Type,
		&pet.Weight,
		&pet.Diseases,
		&pet.Owner.UserID,
		&pet.PetPhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting pet by id. Err msg:%v.", err)
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
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occured while deleting pet. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Pet deleted, rows affectet: %d", result)
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
		p.Owner.UserID,
		p.PetPhotoURL,
		p.PetID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating pet. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Updated pet with id = %d,rows affectet: %d ", p.PetID, result)
	return nil
}

// IsPetValid ...
func (r *PetRepository) IsPetValid(id int) (bool, error) {
	var exist bool
	err := r.Store.Db.QueryRow(`SELECT EXISTS (SELECT id FROM pet WHERE id = $1))`,
		id).Scan(&exist)
	if err != nil {
		return false, err
	}

	if exist {
		return true, nil
	}

	return false, fmt.Errorf("Pet with id = %d does not exist", id)
}
