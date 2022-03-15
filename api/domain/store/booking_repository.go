package store

import (
	"fmt"
	"goReact/domain/model"
	"time"
)

// BookingRepository ...
type BookingRepository struct {
	Store *Store
}

// Create booking and save it to DB
func (r *BookingRepository) Create(b *model.BookingDTO) (*model.Booking, error) {
	if err := r.Store.Db.QueryRow(
		`INSERT INTO booking
		(seat_id, pet_id, employee_id, status, start_date, end_date, paid, notes, transactionId)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		b.SeatID,
		b.PetID,
		b.EmployeeID,
		string(b.Status),
		b.StartDate,
		b.EndDate,
		b.Paid,
		b.Notes,
		b.TransactionID,
	).Scan(&b.BookingID); err != nil {
		r.Store.Logger.Errorf("Eror occured while creating booking. Err msg: %w", err)
		return nil, err
	}

	booking, err := r.ModelFromDTO(b)
	if err != nil {
		return nil, err
	}

	r.Store.Logger.Infof("Booking with id %d was created.", b.BookingID)

	return booking, nil
}

// GetAll returns all bookings
func (r *BookingRepository) GetAll() (*[]model.BookingDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM booking")
	if err != nil {
		r.Store.Logger.Errorf("Eror occured while getting all bookings. Err msg: %w", err)
		return nil, err
	}
	bookings := []model.BookingDTO{}

	for rows.Next() {
		booking := model.BookingDTO{}
		err := rows.Scan(
			&booking.BookingID,
			&booking.SeatID,
			&booking.PetID,
			&booking.EmployeeID,
			&booking.Status,
			&booking.StartDate,
			&booking.EndDate,
			&booking.Notes,
			&booking.Paid,
			&booking.TransactionID,
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
	bookingDTO := &model.BookingDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM booking WHERE id = $1",
		id).Scan(
		&bookingDTO.BookingID,
		&bookingDTO.SeatID,
		&bookingDTO.PetID,
		&bookingDTO.EmployeeID,
		&bookingDTO.Status,
		&bookingDTO.StartDate,
		&bookingDTO.EndDate,
		&bookingDTO.Notes,
		&bookingDTO.Paid,
		&bookingDTO.TransactionID,
	); err != nil {
		r.Store.Logger.Errorf("Eror occured while searching booking. Err msg: %w", err)
		return nil, err
	}

	booking, err := r.ModelFromDTO(bookingDTO)
	if err != nil {
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

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting booking. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while deleting booking. Err msg:%v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Booking with id %d was deleted", id)
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
	transactionID := "transactionId"
	if b.TransactionID != 0 {
		transactionID = fmt.Sprintf("%d", b.TransactionID)
	}

	result, err := r.Store.Db.Exec(
		fmt.Sprintf(`UPDATE booking SET
		seat_id = %s, pet_id = %s, employee_id = %s, status = %s, start_date = %s, end_date = %s, paid = %s, notes = %s, transactionId = %s
		WHERE id = $1`,
			seatID,
			petID,
			employeeID,
			status,
			startDate,
			endDate,
			paid,
			notes,
			transactionID,
		), b.BookingID)
	if err != nil {
		r.Store.Logger.Errorf("Erorr occured while updating booking. Err msg: %w", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating booking. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating booking. Err msg:%v.", err)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Booking with id %d was updated", b.BookingID)
	return nil
}

// ModelFromDTO ...
func (r *BookingRepository) ModelFromDTO(dto *model.BookingDTO) (*model.Booking, error) {
	seat, err := r.Store.Seat().FindByID(dto.SeatID)
	if err != nil {
		return nil, err
	}
	pet, err := r.Store.Pet().FindByID(dto.PetID)
	if err != nil {
		return nil, err
	}
	employee, err := r.Store.Employee().FindByID(dto.EmployeeID)
	if err != nil {
		return nil, err
	}

	return &model.Booking{
		BookingID:     dto.BookingID,
		Seat:          *seat,
		Pet:           *pet,
		Employee:      *employee,
		Status:        model.BookingStatus(dto.Status),
		StartDate:     dto.StartDate,
		EndDate:       dto.EndDate,
		Notes:         dto.Notes,
		TransactionID: dto.TransactionID,
		Paid:          dto.Paid,
	}, nil
}
