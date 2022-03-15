package store

import (
	"errors"
	"hotel/domain/model"
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
		r.Store.Logger.Errorf("Error occured while creating seat. Err msg:%v.", err)
		return nil, err
	}
	return s, nil
}

// GetAll returns all seats
func (r *SeatRepository) GetAll() (*[]model.SeatDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM seat")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all seats. Err msg: %v", err)
	}
	seats := []model.SeatDTO{}

	for rows.Next() {
		seat := model.SeatDTO{}
		err := rows.Scan(
			&seat.SeatID,
			&seat.RoomID,
			&seat.Description,
			&seat.RentFrom,
			&seat.RentTo,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all seats. Err msg: %v", err)
			continue
		}
		seats = append(seats, seat)
	}
	return &seats, nil
}

//FindByID searchs and returns seat by ID
func (r *SeatRepository) FindByID(id int) (*model.SeatDTO, error) {
	seatDTO := &model.SeatDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM seat WHERE id = $1",
		id).Scan(
		&seatDTO.SeatID,
		&seatDTO.RoomID,
		&seatDTO.Description,
		&seatDTO.RentFrom,
		&seatDTO.RentTo,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting seat by id. Err msg:%v.", err)
		return nil, err
	}
	return seatDTO, nil
}

// Delete seat from DB by ID
func (r *SeatRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM seat WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg:%v.", err)
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
		r.Store.Logger.Errorf("Error occured while updating seat. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Updated seat with id = %d,rows affectet: %d ", s.SeatID, result)
	return nil
}

// SeatFromDTO ...
func (r *SeatRepository) SeatFromDTO(dto *model.SeatDTO) (*model.Seat, error) {
	roomDTO, err := r.Store.Room().FindByID(dto.RoomID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while convetring roomDTO. Err msg: %v", err)
		return nil, err
	}

	room, err := r.Store.Room().RoomFromDTO(roomDTO)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while convetring roomDTO. Err msg: %v", err)
		return nil, err
	}

	return &model.Seat{
		SeatID:      dto.RoomID,
		Description: dto.Description,
		Room:        *room,
		RentFrom:    dto.RentFrom,
		RentTo:      dto.RentTo,
	}, nil
}
