package store

import (
	"admin/domain/model"
	"errors"
	"log"
)

// HotelRepository ...
type HotelRepository struct {
	Store *Store
}

// Create hotel and save it to DB
func (r *HotelRepository) Create(h *model.Hotel) (*model.Hotel, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO hotel (name, address) VALUES ($1, $2) RETURNING id", h.Name, h.Address).Scan(&h.HotelID); err != nil {

		return nil, err
	}
	return h, nil
}

// GetAll returns all hotels
func (r *HotelRepository) GetAll() (*[]model.Hotel, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM hotel")
	if err != nil {
		log.Print(err)
	}
	hotels := []model.Hotel{}

	for rows.Next() {
		hotel := model.Hotel{}
		err := rows.Scan(
			&hotel.HotelID,
			&hotel.Name,
			&hotel.Address,
		)
		if err != nil {
			log.Print(err)
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
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return hotel, nil
}

// Delete hotel from DB by ID
func (r *HotelRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM hotel WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("no rows affected")
	}
	log.Printf("Hotel deleted, rows affectet: %d", result)
	return nil
}

// Update hotel from DB
func (r *HotelRepository) Update(e *model.Hotel) error {

	result, err := r.Store.Db.Exec(
		"UPDATE hotel SET name = $1, address = $2 WHERE id = $3",
		e.Name,
		e.Address,
		e.HotelID,
	)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Hotel updated, rows affectet: %d", result)
	return nil
}
