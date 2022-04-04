package store

import (
	"admin/domain/model"
	"errors"
)

// BookingRepository ...
type BookingRepository struct {
	Store *Store
}

// Create booking and save it to DB
func (r *BookingRepository) Create(b *model.Booking) (*model.Booking, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO booking (seat_id, pet_id, employee_id, status, start_date, end_date, notes , paid) VALUES ($1, $2, $3, $4, $5, $6, $7 , $8) RETURNING id",
		b.Seat.SeatID,
		b.Pet.PetID,
		b.Employee.EmployeeID,
		string(b.Status),
		b.StartDate,
		b.EndDate,
		b.Notes,
		b.Paid,
	).Scan(&b.BookingID); err != nil {
		r.Store.Logger.Errorf("Error occurred while creating booking. Err msg:%v.", err)
		return nil, err
	}
	return b, nil
}

// GetAll returns all bookings
func (r *BookingRepository) GetAll() (*[]model.Booking, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM booking")
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while getting all bookings. Err msg: %v", err)
		return nil, err
	}
	bookings := []model.Booking{}

	for rows.Next() {
		booking := model.Booking{}
		err := rows.Scan(
			&booking.BookingID,
			&booking.Seat.SeatID,
			&booking.Pet.PetID,
			&booking.Employee.EmployeeID,
			&booking.Status,
			&booking.StartDate,
			&booking.EndDate,
			&booking.Notes,
			&booking.Paid,
		)
		if err != nil {
			r.Store.Logger.Errorf("Error occurred while getting all bookings. Err msg: %v", err)
			continue
		}
		bookings = append(bookings, booking)
	}
	return &bookings, nil
}

//FindByID searchs and returns booking by ID
func (r *BookingRepository) FindByID(id int) (*model.Booking, error) {
	booking := &model.Booking{}
	if err := r.Store.Db.QueryRow("SELECT * FROM booking WHERE id = $1",
		id).Scan(
		&booking.BookingID,
		&booking.Seat.SeatID,
		&booking.Pet.PetID,
		&booking.Employee.EmployeeID,
		&booking.Status,
		&booking.StartDate,
		&booking.EndDate,
		&booking.Notes,
		&booking.Paid,
	); err != nil {
		r.Store.Logger.Errorf("Error occurred while getting booking by id. Err msg: %v", err)
		return nil, err
	}
	return booking, nil
}

// Delete booking from DB by ID
func (r *BookingRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM booking WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while deleting booking. Err msg: %w", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while deleting booking. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("no rows affected")
		r.Store.Logger.Errorf("Error occurred while deleting booking. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Errorf("Booking deleted, rows affectet: %d", result)
	return nil
}

// Update booking from DB
func (r *BookingRepository) Update(b *model.Booking) error {

	result, err := r.Store.Db.Exec(
		"UPDATE booking SET seat_id = $1, pet_id = $2, employee_id = $3, status = $4, start_date = $5, end_date = $6, notes = $7, paid = $8 WHERE id = $9",
		b.Seat.SeatID,
		b.Pet.PetID,
		b.Employee.EmployeeID,
		string(b.Status),
		b.StartDate,
		b.EndDate,
		b.Notes,
		b.Paid,
		b.BookingID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occurred while updating booking. Err msg:%v.", err)
		return err
	}
	r.Store.Logger.Errorf("Booking updated, rows affectet: %d", result)
	return nil
}
