package store

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/server/handler/pagination"
)

// RoomRepository ...
type RoomRepository struct {
	Store *Store
}

// Create room and save it to DB
func (r *RoomRepository) Create(room *model.Room) (*int, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO room (pet_type, number, hotel_id , photoUrl) VALUES ($1, $2, $3, $4) RETURNING id",
		room.PetType,
		room.RoomNumber,
		room.Hotel.HotelID,
		room.PhotoURL,
	).Scan(&room.RoomID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating room. Err msg: %v.", err)
		return nil, err
	}

	r.Store.Logger.Infof("Room with id %d was created.", room.RoomID)

	return &room.RoomID, nil
}

// GetAll returns all rooms
func (r *RoomRepository) GetAll() (*[]model.RoomDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM room")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %v", err)
		return nil, err
	}
	rooms := []model.RoomDTO{}

	for rows.Next() {
		room := model.RoomDTO{}
		err := rows.Scan(
			&room.RoomID,
			&room.RoomNumber,
			&room.PetType,
			&room.HotelID,
			&room.PhotoURL,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while getting all rooms. Err msg: %v", err)
			continue
		}
		rooms = append(rooms, room)
	}
	return &rooms, nil
}

//FindByID searchs and returns room by ID
func (r *RoomRepository) FindByID(id int) (*model.RoomDTO, error) {
	roomDTO := &model.RoomDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM room WHERE id = $1",
		id).Scan(
		&roomDTO.RoomID,
		&roomDTO.RoomNumber,
		&roomDTO.PetType,
		&roomDTO.HotelID,
		&roomDTO.PhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting room by id. Err msg: %v.", err)
		return nil, err
	}

	return roomDTO, nil
}

// Delete room from DB by ID
func (r *RoomRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM room WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Room with id %d was deleted", id)
	return nil
}

// Update room from DB
func (r *RoomRepository) Update(rm *model.Room) error {

	number := "number"
	if rm.RoomNumber != 0 {
		number = fmt.Sprintf("%d", rm.RoomNumber)
	}
	petType := "pet_type"
	if rm.PetType != "" {
		petType = fmt.Sprintf("'%s'", string(rm.PetType))
	}
	hotelID := "hotel_id"
	if rm.Hotel.HotelID != 0 {
		hotelID = fmt.Sprintf("%d", rm.Hotel.HotelID)
	}
	photoURL := "photoUrl"
	if rm.PhotoURL != "" {
		photoURL = fmt.Sprintf("'%s'", rm.PhotoURL)
	}
	result, err := r.Store.Db.Exec(fmt.Sprintf(
		`UPDATE room SET 
		number = %s, pet_type = %s, hotel_id = %s, photoUrl = %s 
		WHERE id = $1`,
		number,
		petType,
		hotelID,
		photoURL,
	), rm.RoomID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg: %v.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Room with id %d was updated", rm.RoomID)
	return nil
}

// GetAllPagination returns all rooms with limit and offset (limit - max cout off rows on the page)
//offset means which row are first
func (r *RoomRepository) GetAllPagination(p *pagination.Page) (*[]model.RoomDTO, error) {
	p.CalculateOffset()
	rows, err := r.Store.Db.Query("SELECT * FROM ROOM OFFSET $1 LiMIT $2", p.Offset, p.PageSize)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting rooms. Err msg: %v", err)
		return nil, err
	}
	rooms := []model.RoomDTO{}

	for rows.Next() {
		room := model.RoomDTO{}
		err := rows.Scan(
			&room.RoomID,
			&room.RoomNumber,
			&room.PetType,
			&room.HotelID,
			&room.PhotoURL,
		)
		if err != nil {
			r.Store.Logger.Debugf("Error occured while scaning rooms. Err msg: %v", err)
			continue
		}
		rooms = append(rooms, room)
	}

	return &rooms, nil
}

// GetTotalRows ...
func (r *RoomRepository) GetTotalRows() (int, error) {
	var c int
	err := r.Store.Db.QueryRow("SELECT COUNT(*) FROM ROOM").Scan(&c)
	if err != nil {
		r.Store.Logger.Errorf("Error occured counting  rooms. Err msg: %v", err)
		return 0, err
	}

	return c, nil
}

// ModelFromDTO ...
func (r *RoomRepository) ModelFromDTO(dto *model.RoomDTO) (*model.Room, error) {
	hotel, err := r.Store.Hotel().FindByID(dto.HotelID)
	if err != nil {
		return nil, err
	}

	return &model.Room{
		RoomID:     dto.RoomID,
		RoomNumber: dto.RoomNumber,
		PetType:    model.PetType(dto.PetType),
		Hotel:      *hotel,
		PhotoURL:   dto.PhotoURL,
	}, nil
}
