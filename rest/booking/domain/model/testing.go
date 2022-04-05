package model

import "time"

// TestBookingDTO ...
func TestBookingDTO() *BookingDTO {
	a := time.Now().AddDate(0, 0, 1)
	b := time.Now().AddDate(0, 0, 10)
	paid := true
	return &BookingDTO{
		BookingID:     1,
		SeatID:        1,
		PetID:         1,
		EmployeeID:    1,
		Status:        string(BookingStatusInProgress),
		StartDate:     &a,
		EndDate:       &b,
		Notes:         "Notes",
		TransactionID: 1,
		Paid:          &paid,
	}
}

// TestBooking ...
func TestBooking() *Booking {
	a := time.Now().AddDate(0, 0, 1)
	b := time.Now().AddDate(0, 0, 10)
	paid := true
	return &Booking{
		BookingID:     1,
		SeatID:        1,
		PetID:         1,
		EmployeeID:    1,
		Status:        BookingStatusInProgress,
		StartDate:     &a,
		EndDate:       &b,
		Notes:         "Notes",
		TransactionID: 1,
		Paid:          &paid,
	}
}
