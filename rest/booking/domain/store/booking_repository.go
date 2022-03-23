package store

import (
	"booking/domain/model"
	"fmt"
	"time"
)

// BookingRepository ...
type BookingRepository struct {
	Store *Store
}

// Create booking and save it to DB
func (r *BookingRepository) Create(b *model.Booking) (*model.Booking, error) {
	if err := r.Store.Db.QueryRow(
		`INSERT INTO booking
		(seat_id, pet_id, employee_id, transactionId, status, start_date, end_date, paid, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, &9) RETURNING id`,
		b.SeatID,
		b.PetID,
		b.EmployeeID,
		b.TransactionID,
		string(b.Status),
		b.StartDate,
		b.EndDate,
		b.Paid,
		b.Notes,
	).Scan(&b.BookingID); err != nil {
		r.Store.Logger.Errorf("Eror occured while creating booking. Err msg: %w", err)
		return nil, err
	}
	return b, nil
}

// GetAll returns all bookings
func (r *BookingRepository) GetAll() (*[]model.Booking, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM booking")
	if err != nil {
		r.Store.Logger.Errorf("Eror occured while getting all bookings. Err msg: %w", err)
		return nil, err
	}
	bookings := []model.Booking{}

	for rows.Next() {
		booking := model.Booking{}
		err := rows.Scan(
			&booking.BookingID,
			&booking.SeatID,
			&booking.PetID,
			&booking.EmployeeID,
			&booking.TransactionID,
			&booking.Status,
			&booking.StartDate,
			&booking.EndDate,
			&booking.Paid,
			&booking.Notes,
		)
		if err != nil {
			r.Store.Logger.Errorf("Eror occured while getting all bookings. Err msg: %w", err)
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
		&booking.SeatID,
		&booking.PetID,
		&booking.EmployeeID,
		&booking.TransactionID,
		&booking.Status,
		&booking.StartDate,
		&booking.EndDate,
		&booking.Paid,
		&booking.Notes,
	); err != nil {
		r.Store.Logger.Errorf("Eror occured while searching booking. Err msg: %w", err)
		return nil, err
	}
	return booking, nil
}

// Delete booking from DB by ID
func (r *BookingRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM booking WHERE id = $1", id)
	if err != nil {
		r.Store.Logger.Errorf("Eror occured while deleting booking. Err msg: %w", err)
		return err
	}
	r.Store.Logger.Errorf("Booking deleted, rows affectet: %d", result)
	return nil
}

// Update booking from DB
func (r *BookingRepository) Update(b *model.Booking) error {
	seatID := "seat_id"
	if b.SeatID != 0 {
		seatID = fmt.Sprintf("%d", b.SeatID)
	}
	petID := "pet_id"
	if b.PetID != 0 {
		petID = fmt.Sprintf("%d", b.PetID)
	}
	employeeID := "employee_id"
	if b.EmployeeID != 0 {
		employeeID = fmt.Sprintf("%d", b.EmployeeID)
	}
	transactionID := "transactionId"
	if b.TransactionID != 0 {
		transactionID = fmt.Sprintf("%d", b.TransactionID)
	}
	status := "status"
	if b.Status != "" {
		status = fmt.Sprintf("'%s'", string(b.Status))
	}
	startDate := "start_date"
	if b.StartDate != nil {
		startDate = fmt.Sprintf("'%s'", b.StartDate.Format(time.RFC3339))
	}
	endDate := "end_date"
	if b.EndDate != nil {
		endDate = fmt.Sprintf("'%s'", b.EndDate.Format(time.RFC3339))
	}
	paid := "paid"
	if b.Paid != nil {
		paid = fmt.Sprintf("%v", *b.Paid)
	}
	notes := "notes"
	if b.Notes != "" {
		notes = fmt.Sprintf("'%s'", b.Notes)
	}

	result, err := r.Store.Db.Exec(
		fmt.Sprintf(`UPDATE booking SET
		seat_id = %s, pet_id = %s, employee_id = %s, transactionId = %s, status = %s, start_date = %s, end_date = %s, paid = %s, notes = %s
		WHERE id = $1`,
			seatID,
			petID,
			employeeID,
			transactionID,
			status,
			startDate,
			endDate,
			paid,
			notes,
		), b.BookingID)
	if err != nil {
		r.Store.Logger.Errorf("Eror occured while updating booking. Err msg: %w", err)
		return err
	}
	r.Store.Logger.Errorf("Booking updated, rows affectet: %d", result)
	return nil
}
