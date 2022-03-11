package store

import (
	"goReact/domain/model"
	"goReact/webapp/server/handler/pagination"
	"log"
)

// RoomRepository ...
type RoomRepository struct {
	Store *Store
}

// Create room and save it to DB
func (r *RoomRepository) Create(rDTO *model.RoomDTO) (*model.Room, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO room (pet_type, number, hotel_id , photo) VALUES ($1, $2, $3, $4) RETURNING id",
		rDTO.PetType,
		rDTO.RoomNumber,
		rDTO.HotelID,
		rDTO.RoomPhotoURL,
	).Scan(&rDTO.RoomID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating room. Err msg: %w.", err)
		return nil, err
	}

	room, err := r.ModelFromDTO(rDTO)
	if err != nil {
		return nil, err
	}

	r.Store.Logger.Infof("Room with id %d was created.", room.RoomID)

	return room, nil
}

// GetAll returns all rooms
func (r *RoomRepository) GetAll() (*[]model.RoomDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM room")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %w", err)
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
			r.Store.Logger.Errorf("Error occured while getting all rooms. Err msg: %w", err)
			continue
		}
		rooms = append(rooms, room)
	}
	return &rooms, nil
}

//FindByID searchs and returns room by ID
func (r *RoomRepository) FindByID(id int) (*model.Room, error) {
	roomDTO := &model.RoomDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM room WHERE id = $1",
		id).Scan(
		&roomDTO.RoomID,
		&roomDTO.RoomNumber,
		&roomDTO.PetType,
		&roomDTO.HotelID,
		&roomDTO.RoomPhotoURL,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while getting room by id. Err msg: %w.", err)
		return nil, err
	}

	room, err := r.ModelFromDTO(roomDTO)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// Delete room from DB by ID
func (r *RoomRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM room WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %w.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %w.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting room. Err msg: %w.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Room with id %d was deleted", id)
	return nil
}

// Update room from DB
func (r *RoomRepository) Update(rm *model.RoomDTO) error {
	result, err := r.Store.Db.Exec(
		"UPDATE room SET number = $1, pet_type = $2, hotel_id = $3, photo = $4 WHERE id = $5",
		rm.RoomNumber,
		rm.PetType,
		rm.HotelID,
		rm.RoomPhotoURL,
		rm.RoomID)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg: %w.", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating room. Err msg:%v.", err)
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
		log.Print(err)
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
		)
		if err != nil {
			log.Print(err)
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
		log.Print(err.Error())
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
		RoomID:       dto.RoomID,
		RoomNumber:   dto.RoomNumber,
		PetType:      model.PetType(dto.PetType),
		Hotel:        *hotel,
		RoomPhotoURL: dto.RoomPhotoURL,
	}, nil
}
