package store

import (
	"goReact/domain/model"

	"github.com/lib/pq"
)

// HotelRepository ...
type HotelRepository struct {
	Store *Store
}

// Create hotel and save it to DB
func (r *HotelRepository) Create(h *model.Hotel) (*model.Hotel, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO hotel (name, address, coordinates ) VALUES ($1, $2 , $3) RETURNING id",
		h.Name,
		h.Address,
		pq.Array(h.Coordinates)).Scan(&h.HotelID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating hotel. Err msg:%v.", err)
		return nil, err
	}

	r.Store.Logger.Info("Created hotel with id = %d", h.HotelID)
	return h, nil
}

// GetAll returns all hotels
func (r *HotelRepository) GetAll() (*[]model.Hotel, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM hotel")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all hotels. Err msg: %v", err)
	}
	hotels := []model.Hotel{}

	for rows.Next() {
		hotel := model.Hotel{}
		err := rows.Scan(
			&hotel.HotelID,
			&hotel.Name,
			&hotel.Address,
			pq.Array(&hotel.Coordinates),
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all hotels. Err msg: %v", err)
			continue
		}
		hotels = append(hotels, hotel)
	}
	return &hotels, nil
}

//FindByID searchs and returns hotel by ID
func (r *HotelRepository) FindByID(id int) (*model.Hotel, error) {
	hotel := &model.Hotel{}
	if err := r.Store.Db.QueryRow("SELECT * FROM hotel WHERE id = $1",
		id).Scan(
		&hotel.HotelID,
		&hotel.Name,
		&hotel.Address,
		pq.Array(&hotel.Coordinates),
	); err != nil {
		r.Store.Logger.Errorf("Error occured while searching hotel by id. Err msg:%v.", err)
		return nil, err
	}
	return hotel, nil
}

// Delete hotel from DB by ID
func (r *HotelRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM hotel WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting hotel. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting hotel. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting hotel. Err msg:%v.", err)
		return ErrNoRowsAffected
	}
	r.Store.Logger.Info("Hotel with id %d was deleted.", id)
	return nil
}

// Update hotel from DB
func (r *HotelRepository) Update(h *model.Hotel) error {
	result, err := r.Store.Db.Exec(
		"UPDATE hotel SET name = $1, address = $2 , coordinates = $3 WHERE id = $4",
		h.Name,
		h.Address,
		pq.Array(h.Coordinates),
		h.HotelID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating hotel. Err msg:%v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating hotel. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating hotel. Err msg:%v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Info("Hotel with id = %d was updated", h.HotelID)
	return nil
}
