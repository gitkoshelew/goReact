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

var permission_update model.Permission = model.Permission{Name: model.UpdateBooking}

// update booking ...
func UpdateBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_update.Name)
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

		b, err := s.Booking().FindByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while finding booking by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		seatID, err := strconv.Atoi(r.FormValue("SeatID"))
		if err == nil {
			if seatID != 0 {
				seat, err := s.Seat().FindByID(seatID)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error occured while finding seat by id. Err msg:%v. ", err), http.StatusBadRequest)
					return
				}
				b.Seat = *seat
			}

		}

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err == nil {
			if petID != 0 {
				pet, err := s.Pet().FindByID(petID)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error occured while finding pet by id. Err msg:%v. ", err), http.StatusBadRequest)
					return
				}
				b.Pet = *pet
			}

		}

		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err == nil {
			if employeeID != 0 {
				employee, err := s.Employee().FindByID(employeeID)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error occured while finding employee by id. Err msg:%v. ", err), http.StatusBadRequest)
					return
				}
				b.Employee = *employee
			}

		}

		status := r.FormValue("Status")
		if status != "" {
			b.Status = model.BookingStatus(status)
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
			b.StartDate = startDate
		}

		endDate := r.FormValue("EndDate")
		if endDate != "" {
			endDate, err := time.Parse(layout, r.FormValue("EndDate"))
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate"))
				return
			}
			b.EndDate = endDate
		}

		notes := r.FormValue("Notes")
		if notes != "" {
			b.Notes = notes
		}

		paid := r.FormValue("Paid")
		if paid != "" {
			paid, err := strconv.ParseBool(paid)
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("paid")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid"))
				return
			}
			b.Paid = paid
		}

		err = b.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Booking().Update(b)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating booking. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homebookings", http.StatusFound)

	}
}
