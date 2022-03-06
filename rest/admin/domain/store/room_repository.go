package store

import (
	"admin/domain/model"
	"errors"
)

// RoomRepository ...
type RoomRepository struct {
	Store *Store
}

// Create room and save it to DB
func (r *RoomRepository) Create(rm *model.Room) (*model.Room, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO room (pet_type, number, hotel_id) VALUES ($1, $2, $3) RETURNING id",
		string(rm.PetType), rm.RoomNumber, rm.Hotel.HotelID).Scan(&rm.RoomID); err != nil {
		r.Store.Logger.Errorf("Can't create room. Err msg:%v.", err)
		return nil, err
	}
	r.Store.Logger.Info("Created room with id = %d", rm.RoomID)

	return rm, nil
}

// GetAll returns all rooms
func (r *RoomRepository) GetAll() (*[]model.Room, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM room")
	if err != nil {
		r.Store.Logger.Errorf("Can't find rooms. Err msg: %v", err)
	}
	rooms := []model.Room{}

	for rows.Next() {
		room := model.Room{}
		err := rows.Scan(
			&room.RoomID,
			&room.RoomNumber,
			&room.PetType,
			&room.Hotel.HotelID,
			&room.RoomPhotoURL,

		)
		if err != nil {
			r.Store.Logger.Errorf("Can't find rooms. Err msg: %v", err)
			continue
		}
		rooms = append(rooms, room)
	}
	return &rooms, nil
}

//FindByID searchs and returns room by ID
func (r *RoomRepository) FindByID(id int) (*model.Room, error) {
	room := &model.Room{}
	if err := r.Store.Db.QueryRow("SELECT * FROM room WHERE id = $1",
		id).Scan(
		&room.RoomID,
		&room.RoomNumber,
		&room.PetType,
		&room.Hotel.HotelID,
	); err != nil {
		r.Store.Logger.Errorf("Can't find room. Err msg:%v.", err)
		return nil, err
	}
	return room, nil
}

// Delete room from DB by ID
func (r *RoomRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM room WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Can't delete room. Err msg:%v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Can't delete room. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Can't delete room. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Room deleted, rows affectet: %d", result)
	return nil
}

// Update room from DB
func (r *RoomRepository) Update(rm *model.Room) error {

	result, err := r.Store.Db.Exec(
		"UPDATE room SET number = $1, pet_type = $2, hotel_id = $3 WHERE id = $4",
		rm.RoomNumber,
		string(rm.PetType),
		rm.Hotel.HotelID,
		rm.RoomID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Can't update room. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Info("Update room with id = %d,rows affectet: %d ", rm.RoomID, result)
	return nil
}
