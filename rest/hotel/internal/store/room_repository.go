package store

import (
	"errors"
	"hotel/domain/model"
	"hotel/pkg/pagination"
)

// RoomRepository ...
type RoomRepository struct {
	Store *Store
}

// Create room and save it to DB
func (r *RoomRepository) Create(rm *model.Room) (*model.Room, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO room (pet_type, number, hotel_id , photo) VALUES ($1, $2, $3, $4) RETURNING id",
		string(rm.PetType),
		rm.RoomNumber,
		rm.Hotel.HotelID,
		rm.RoomPhotoURL).Scan(&rm.RoomID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating room. Err msg:%v.", err)
		return nil, err
	}
	r.Store.Logger.Info("Created room with id = %d", rm.RoomID)
	return rm, nil
}

// GetAll returns all rooms
func (r *RoomRepository) GetAll() (*[]model.RoomDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM room")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %v", err)
	}
	rooms := []model.RoomDTO{}

	for rows.Next() {
		room := model.RoomDTO{}
		err := rows.Scan(
			&room.RoomID,
			&room.RoomNumber,
			&room.PetType,
			&room.HotelID,
			&room.RoomPhotoURL,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %v", err)
			continue
		}
		rooms = append(rooms, room)
	}
	return &rooms, nil
}

//FindByID searchs and returns room by ID
func (r *RoomRepository) FindByID(id int) (*model.RoomDTO, error) {
	room := &model.RoomDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM room WHERE id = $1",
		id).Scan(
		&room.RoomID,
		&room.RoomNumber,
		&room.PetType,
		&room.HotelID,
		&room.RoomPhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting room by id. Err msg:%v.", err)
		return nil, err
	}
	return room, nil
}

// Delete room from DB by ID
func (r *RoomRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM room WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Room deleted, rows affectet: %d", result)
	return nil
}

// Update room from DB
func (r *RoomRepository) Update(rm *model.Room) error {

	result, err := r.Store.Db.Exec(
		"UPDATE room SET number = $1, pet_type = $2, hotel_id = $3, photo = $4 WHERE id = $5",
		rm.RoomNumber,
		string(rm.PetType),
		rm.Hotel.HotelID,
		rm.RoomPhotoURL,
		rm.RoomID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Updated room with id = %d,rows affectet: %d ", rm.RoomID, result)
	return nil
}

// GetAllPagination returns all rooms with limit and offset (limit - max cout off rows on the page)
//offset means which row are first
func (r *RoomRepository) GetAllPagination(p *pagination.Page) (*[]model.RoomDTO, error) {
	p.CalculateOffset()
	rows, err := r.Store.Db.Query("SELECT * FROM ROOM OFFSET $1 LiMIT $2", p.Offset, p.PageSize)
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
			&room.RoomPhotoURL,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %v", err)
			continue
		}
		rooms = append(rooms, room)
	}

	return &rooms, nil
}

// RoomFromDTO ...
func (r *RoomRepository) RoomFromDTO(dto *model.RoomDTO) (*model.Room, error) {
	r.Store.Logger.Info("dto  ", dto)
	r.Store.Logger.Info("dto.HotelID   ", dto.HotelID)
	id := dto.HotelID
	r.Store.Logger.Info("id   ", id)
	hotel, err := r.Store.Hotel().FindByID(id)
	r.Store.Logger.Info("hotel", hotel)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while converting roomDTO. Err msg: %v", err)
		return nil, err
	}
	return &model.Room{
		RoomID:       dto.RoomID,
		RoomNumber:   dto.RoomNumber,
		PetType:      dto.PetType,
		Hotel:        *hotel,
		RoomPhotoURL: dto.RoomPhotoURL,
	}, nil

}

// GetTotalRows ...
func (r *RoomRepository) GetTotalRows() (int, error) {
	var c int
	err := r.Store.Db.QueryRow("SELECT COUNT(*) FROM ROOM").Scan(&c)
	if err != nil {
		r.Store.Logger.Errorf("Error occured whilegetting total rows. Err msg: %v", err)
		return 0, err
	}

	return c, nil
}
