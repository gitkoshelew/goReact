package store

import (
	"errors"
	"goReact/domain/model"
	"log"
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
		log.Print(err)
		return nil, err
	}
	return rm, nil
}

// GetAll returns all rooms
func (r *RoomRepository) GetAll() (*[]model.Room, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM room")
	if err != nil {
		log.Print(err)
	}
	rooms := []model.Room{}

	for rows.Next() {
		room := model.Room{}
		err := rows.Scan(
			&room.RoomID,
			&room.RoomNumber,
			&room.PetType,
			&room.Hotel.HotelID,
		)
		if err != nil {
			log.Print(err)
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
		log.Printf(err.Error())
		return nil, err
	}
	return room, nil
}

// Delete room from DB by ID
func (r *RoomRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM room WHERE id = $1", id)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("No rows affected")
	}
	log.Printf("Room deleted, rows affectet: %d", result)
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
		log.Printf(err.Error())
		return err
	}
	log.Printf("Room updated, rows affectet: %d", result)
	return nil
}
