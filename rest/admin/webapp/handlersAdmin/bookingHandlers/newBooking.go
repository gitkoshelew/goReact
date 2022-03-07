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

var permission_create model.Permission = model.Permission{Name: model.CreatBooking}

// NewBooking ...
func NewBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_create.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		seatID, err := strconv.Atoi(r.FormValue("SeatID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("SeatID"))
			return
		}
		seat, err := s.Seat().FindByID(seatID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("PetID"))
			return
		}
		pet, err := s.Pet().FindByID(petID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EmployeeID"))
			return
		}
		employee, err := s.Employee().FindByID(employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status := r.FormValue("Status")
		layout := "2006-01-02"
		startDate, err := time.Parse(layout, r.FormValue("StartDate"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate"))
			return
		}
		endDate, err := time.Parse(layout, r.FormValue("EndDate"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate"))
			return
		}
		notes := r.FormValue("Notes")

		paid, err := strconv.ParseBool(r.FormValue("Paid"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid"))
			return
		}

		b := model.Booking{
			BookingID: 0,
			Seat:      *seat,
			Pet:       *pet,
			Employee:  *employee,
			Status:    model.BookingStatus(status),
			StartDate: startDate,
			EndDate:   endDate,
			Notes:     notes,
			Paid:      paid,
		}

		err = b.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Booking().Create(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homebookings/", http.StatusFound)

	}
}
