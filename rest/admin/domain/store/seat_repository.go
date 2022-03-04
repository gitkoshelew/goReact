package store

import (
	"admin/domain/model"
	"errors"
)

// SeatRepository ...
type SeatRepository struct {
	Store *Store
}

// Create seat and save it to DB
func (r *SeatRepository) Create(s *model.Seat) (*model.Seat, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO seat (room_id, rent_from, rent_to, description) VALUES ($1, $2, $3, $4) RETURNING id",
		s.Room.RoomID,
		s.RentFrom,
		s.RentTo,
		s.Description,
	).Scan(&s.SeatID); err != nil {
		r.Store.Logger.Errorf("Can't create seat. Err msg:%v.", err)
		return nil, err
	}
	return s, nil
}

// GetAll returns all seats
func (r *SeatRepository) GetAll() (*[]model.Seat, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM seat")
	if err != nil {
		r.Store.Logger.Errorf("Can't find seats. Err msg: %v", err)
	}
	seats := []model.Seat{}

	for rows.Next() {
		seat := model.Seat{}
		err := rows.Scan(
			&seat.SeatID,
			&seat.Room.RoomID,
			&seat.Description,
			&seat.RentFrom,
			&seat.RentTo,
		)
		if err != nil {
			r.Store.Logger.Errorf("Can't find seats. Err msg: %v", err)
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
		&seat.Description,
		&seat.RentFrom,
		&seat.RentTo,
	); err != nil {
		r.Store.Logger.Errorf("Can't find seat. Err msg:%v.", err)
		return nil, err
	}
	return seat, nil
}

// Delete seat from DB by ID
func (r *SeatRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM seat WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Can't delete seat. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Can't delete seat. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Can't delete seat. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Seat deleted, rows affectet: %d", result)
	return nil
}

// Update seat from DB
func (r *SeatRepository) Update(s *model.Seat) error {

	result, err := r.Store.Db.Exec(
		"UPDATE seat SET room_id = $1, rent_from = $2, rent_to = $3, description = $4 WHERE id = $5",
		s.Room.RoomID,
		s.RentFrom,
		s.RentTo,
		s.Description,
		s.SeatID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Can't update seat. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Update seat with id = %d,rows affectet: %d ", s.SeatID, result)
	return nil
}
