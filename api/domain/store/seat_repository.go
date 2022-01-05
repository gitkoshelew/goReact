package store

import (
	"goReact/domain/model"
	"log"
)

// SeatRepository ...
type SeatRepository struct {
	Store *Store
}

// Create seat and save it to DB
func (r *SeatRepository) Create(s *model.Seat) (*model.Seat, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO seat",
		"(room_id, is_free, description)",
		"VALUES ($1, $2, $3) RETURNING id",
		s.Room.RoomID,
		s.IsFree,
		s.Description,
	).Scan(&s.SeatID); err != nil {
		log.Print(err)
		return nil, err
	}
	return s, nil
}

// GetAll returns all seats
func (r *SeatRepository) GetAll() (*[]model.Seat, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM seat")
	if err != nil {
		log.Print(err)
	}
	seats := []model.Seat{}

	for rows.Next() {
		seat := model.Seat{}
		err := rows.Scan(
			&seat.SeatID,
			&seat.Room.RoomID,
			&seat.IsFree,
			&seat.Description,
		)
		if err != nil {
			log.Print(err)
			continue
		}
		seats = append(seats, seat)
	}
	return &seats, nil
}

//FindByID searchs and returns seat by ID
func (r *SeatRepository) FindByID(id int) (*model.Seat, error) {
	seat := &model.Seat{}
	if err := r.Store.Db.QueryRow("SELECT * FROM seat WHERE id = $1",
		id).Scan(
		&seat.SeatID,
		&seat.Room.RoomID,
		&seat.IsFree,
		&seat.Description,
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return seat, nil
}

// Delete seat from DB by ID
func (r *SeatRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM seat WHERE id = $1", id)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Seat deleted, rows affectet: %d", result)
	return nil
}

// Update seat from DB
func (r *SeatRepository) Update(s *model.Seat) error {

	result, err := r.Store.Db.Exec(
		"UPDATE seat SET",
		"room_id = $1, is_free = $2, description = $3",
		"WHERE id = $4",
		s.Room.RoomID,
		s.IsFree,
		s.Description,
		s.SeatID,
	)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Seat updated, rows affectet: %d", result)
	return nil
}
