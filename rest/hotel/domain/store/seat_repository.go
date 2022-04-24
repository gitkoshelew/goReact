package store

import (
	"errors"
	"fmt"
	"hotel/domain/model"
	"hotel/domain/request"
	"time"
)

// SeatRepository ...
type SeatRepository struct {
	Store *Store
}

// Create seat and save it to DB
func (r *SeatRepository) Create(s *model.Seat) (*int, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO seat (room_id, rent_from, rent_to , price) VALUES ($1, $2, $3, $4) RETURNING seat_id",
		s.Room.RoomID,
		s.RentFrom,
		s.RentTo,
		s.Price,
	).Scan(&s.SeatID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating seat. Err msg: %v.", err)
		return nil, err
	}

	r.Store.Logger.Infof("Seat with id %d was created.", s.SeatID)

	return &s.SeatID, nil
}

// GetAll returns all seats
func (r *SeatRepository) GetAll() (*[]model.SeatDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM seat")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all seats. Err msg: %v", err)
		return nil, err
	}
	seats := []model.SeatDTO{}

	for rows.Next() {
		seat := model.SeatDTO{}
		err := rows.Scan(
			&seat.SeatID,
			&seat.RoomID,
			&seat.RentFrom,
			&seat.RentTo,
			&seat.Price,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while getting all seats. Err msg: %v", err)
			continue
		}
		seats = append(seats, seat)
	}
	return &seats, nil
}

//FindByID searchs and returns seat by ID
func (r *SeatRepository) FindByID(id int) (*model.SeatDTO, error) {
	seatDTO := &model.SeatDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM seat WHERE seat_id = $1",
		id).Scan(
		&seatDTO.SeatID,
		&seatDTO.RoomID,
		&seatDTO.RentFrom,
		&seatDTO.RentTo,
		&seatDTO.Price,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting seat by id. Err msg: %v.", err)
		return nil, err
	}

	return seatDTO, nil
}

// Delete seat from DB by ID
func (r *SeatRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM seat WHERE seat_id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg: %v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg: %v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting seat. Err msg: %v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Seat with id %d was deleted", id)
	return nil
}

// Update seat from DB
func (r *SeatRepository) Update(s *model.Seat) error {
	roomID := "room_id"
	if s.Room.RoomID != 0 {
		roomID = fmt.Sprintf("%d", s.Room.RoomID)
	}
	rentFrom := "rent_from"
	if s.RentFrom != nil {
		rentFrom = fmt.Sprintf("'%s'", s.RentFrom.Format(time.RFC3339))
	}
	rentTo := "rent_to"
	if s.RentTo != nil {
		rentTo = fmt.Sprintf("'%s'", s.RentTo.Format(time.RFC3339))
	}

	price := "price"
	if s.Price != 0 {
		price = fmt.Sprintf("%f", s.Price)
	}
	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE seat SET 
		room_id = %s, rent_from = %s, rent_to = %s, price = %s
		WHERE seat_id = $1`,
		roomID,
		rentFrom,
		rentTo,
		price,
	), s.SeatID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating seat. Err msg: %v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating seat. Err msg: %v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating seat. Err msg: %v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Seat with id %d was updated", s.SeatID)

	return nil
}

// ModelFromDTO ...
func (r *SeatRepository) ModelFromDTO(dto *model.SeatDTO) (*model.Seat, error) {
	roomDTO, err := r.Store.Room().FindByID(dto.RoomID)
	if err != nil {
		return nil, err
	}
	room, err := r.Store.Room().ModelFromDTO(roomDTO)
	if err != nil {
		return nil, err
	}

	return &model.Seat{
		SeatID:   dto.RoomID,
		Room:     *room,
		RentFrom: dto.RentFrom,
		RentTo:   dto.RentTo,
		Price:    dto.Price,
	}, nil
}

// FreeSeatsSearching searching free seats by hotel ID, pet type, rentTo and rentFrom data
func (r *SeatRepository) FreeSeatsSearching(req *request.FreeSeatsSearching) (*[]model.SeatDTO, error) {
	rows, err := r.Store.Db.Query(`SELECT * FROM seat 
	WHERE 
	(room_id = any(SELECT seat_id FROM room WHERE pet_type = $1 AND hotel_id = $2)) 
	and 
	($3 >= rent_to or $4 <= rent_from)`,
		req.PetType,
		req.HotelID,
		req.RentFrom,
		req.RentTo)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all seats. Err msg: %v", err)
		return nil, err
	}

	seats := []model.SeatDTO{}

	for rows.Next() {
		seat := model.SeatDTO{}
		err := rows.Scan(
			&seat.SeatID,
			&seat.RoomID,
			&seat.RentFrom,
			&seat.RentTo,
			&seat.Price,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while getting all seats. Err msg: %v", err)
			continue
		}
		seats = append(seats, seat)
	}

	if len(seats) < 1 {
		r.Store.Logger.Debugf("no free seats for data: %v", req)
		return nil, errors.New("no seats available for chosen data")

	}

	r.Store.Logger.Infof("seat IDs = %v", seats)

	return &seats, nil
}
