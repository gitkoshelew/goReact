package bookinghandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permission_update model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "update_booking",
	Descriptoin:  "ability to update a booking"}

// update booking ...
func UpdateBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_update.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("BookingID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("BookingID"))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}
		b, err := s.Booking().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			return
		}

		seatID, err := strconv.Atoi(r.FormValue("SeatID"))
		if err != nil {
			if seatID != 0 {
				seat, err := s.Seat().FindByID(seatID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					s.Logger.Errorf("Cant find seat. Err msg:%v.", err)
					return
				}
				b.Seat = *seat
			}

		}

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			if petID != 0 {
				pet, err := s.Pet().FindByID(petID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					s.Logger.Errorf("Cant find pet. Err msg:%v.", err)
					return
				}
				b.Pet = *pet
			}

		}

		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err != nil {
			if employeeID != 0 {
				employee, err := s.Employee().FindByID(employeeID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					s.Logger.Errorf("Cant find employee. Err msg:%v.", err)
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
				w.WriteHeader(http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate"))
				return
			}
			b.StartDate = startDate
		}

		endDate := r.FormValue("EndDate")
		if endDate != "" {
			endDate, err := time.Parse(layout, r.FormValue("EndDate"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
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
				w.WriteHeader(http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified"))
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

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v.", err)
			return
		}

		err = s.Booking().Update(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't update booking. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Update user with id = %d", b.BookingID)
		http.Redirect(w, r, "/admin/homebookings", http.StatusFound)

	}
}
