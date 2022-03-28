package bookinghandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateBooking}

// UpdateBooking ...
func UpdateBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionUpdate.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("BookingID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("BookingID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("BookingID"))
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bookingDTO, err := s.Booking().FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while finding booking by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		seatID, err := strconv.Atoi(r.FormValue("SeatID"))
		if err == nil {
			if seatID != 0 {
				bookingDTO.SeatID = seatID
			}

		}

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err == nil {
			if petID != 0 {
				bookingDTO.PetID = petID
			}

		}

		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err == nil {
			if employeeID != 0 {
				bookingDTO.EmployeeID = employeeID
			}

		}

		status := r.FormValue("Status")
		if status != "" {
			bookingDTO.Status = status
		}

		layout := "2006-01-02"
		startDate := r.FormValue("StartDate")
		if startDate != "" {
			startDate, err := time.Parse(layout, r.FormValue("StartDate"))
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate"))
				return
			}
			bookingDTO.StartDate = &startDate
		}

		endDate := r.FormValue("EndDate")
		if endDate != "" {
			endDate, err := time.Parse(layout, r.FormValue("EndDate"))
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate"))
				return
			}
			bookingDTO.EndDate = &endDate
		}

		notes := r.FormValue("Notes")
		if notes != "" {
			bookingDTO.Notes = notes
		}

		paid := r.FormValue("Paid")
		if paid != "" {
			paid, err := strconv.ParseBool(paid)
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid"))
				return
			}
			bookingDTO.Paid = &paid
		}

		transactionID, err := strconv.Atoi(r.FormValue("TransactionID"))
		if err == nil {
			if employeeID != 0 {
				bookingDTO.TransactionID = transactionID
			}
		}

		err = bookingDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}
		booking, err := s.Booking().ModelFromDTO(bookingDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		err = s.Booking().Update(booking)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating booking. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homebookings", http.StatusFound)

	}
}
