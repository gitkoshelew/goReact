package store

import (
	"fmt"
	"goReact/domain/model"
	reqandresp "goReact/domain/reqAndResp"
	"time"
)

// SeatRepository ...
type SeatRepository struct {
	Store *Store
}

// Create seat and save it to DB
func (r *SeatRepository) Create(s *model.Seat) (*int, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO seat room_id, description, price VALUES ($1, $2, $3) RETURNING id",
		s.Room.RoomID,
		s.Description,
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
			&seat.Description,
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
	if err := r.Store.Db.QueryRow("SELECT * FROM seat WHERE id = $1",
		id).Scan(
		&seatDTO.SeatID,
		&seatDTO.RoomID,
		&seatDTO.Description,
		&seatDTO.Price,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting seat by id. Err msg: %v.", err)
		return nil, err
	}

	return seatDTO, nil
}

// Delete seat from DB by ID
func (r *SeatRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM seat WHERE id = $1", id)
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
	description := "description"
	if s.Description != "" {
		description = fmt.Sprintf("'%s'", s.Description)
	}

	price := "price"
	if s.Price != 0 {
		price = fmt.Sprintf("%f", s.Price)
	}
	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE seat SET 
		room_id = %s, description = %s, price = %s 
		WHERE id = $1`,
		roomID,
		description,
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
		SeatID:      dto.RoomID,
		Description: dto.Description,
		Room:        *room,
		Price:       dto.Price,
	}, nil
}

// FreeSeatsSearching searching free seats by hotel ID, pet type, rentTo and rentFrom data
func (r *SeatRepository) FreeSeatsSearching(req *reqandresp.FreeSeatsSearching) (*map[int][]int, error) {
	rows, err := r.Store.Db.Query(`SELECT  S.seat_id, start_date, end_date   FROM seat AS S
	LEFT JOIN booking AS B ON(S.seat_id = B.seat_id)
	JOIN room AS R ON (pet_type = $1 AND hotel_id = $2 and S.room_id = R.room_id)
	GROUP BY S.seat_id, start_date, end_date
	ORDER BY start_date`,
		req.PetType,
		req.HotelID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting room info. Err msg: %v", err)
		return nil, err
	}

	roomInfos := []reqandresp.RoomInfo{}

	for rows.Next() {
		roomInfo := reqandresp.RoomInfo{}
		err := rows.Scan(
			&roomInfo.SeatID,
			&roomInfo.StartDate,
			&roomInfo.EndDate,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while getting room info. Err msg: %v", err)
			continue
		}
		roomInfos = append(roomInfos, roomInfo)
	}

	if len(roomInfos) < 1 {
		r.Store.Logger.Debugf("no free seats for data: %v", req)
		return nil, ErrNoFreeSeatsForCurrentRequest
	}

	BookingCalendar := make(map[int][]int)

	for i := 1; i < 31; i++ {
		var seatsIDs []int
		date := time.Now().AddDate(0, 0, i)
		existingSeats := make(map[int]bool)
		for _, d := range roomInfos {
			if d.StartDate == nil && d.EndDate == nil {
				if _, ok := existingSeats[d.SeatID]; !ok {
					existingSeats[d.SeatID] = true
					seatsIDs = append(seatsIDs, d.SeatID)
				}
			} else {
				if date.Before(*d.StartDate) || date.After(*d.EndDate) {
					if _, ok := existingSeats[d.SeatID]; !ok {
						existingSeats[d.SeatID] = true
						seatsIDs = append(seatsIDs, d.SeatID)
					}
				}
			}
		}

		BookingCalendar[i] = seatsIDs
	}

	return &BookingCalendar, nil
}
