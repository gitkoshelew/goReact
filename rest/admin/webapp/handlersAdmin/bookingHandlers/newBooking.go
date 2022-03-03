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

var permission_create model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "create_bookings",
	Descriptoin:  "ability to create a booking"}

// NewBooking ...
func NewBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_create.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}
		seatID, err := strconv.Atoi(r.FormValue("SeatID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			return
		}
		seat, err := s.Seat().FindByID(seatID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find seat. Err msg:%v.", err)
			return
		}

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			return
		}
		pet, err := s.Pet().FindByID(petID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find pet. Err msg:%v.", err)
			return
		}
		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			return
		}
		employee, err := s.Employee().FindByID(employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find employee. Err msg:%v.", err)
			return
		}
		status := r.FormValue("Status")
		layout := "2006-01-02"
		startDate, err := time.Parse(layout, r.FormValue("StartDate"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
			return
		}
		endDate, err := time.Parse(layout, r.FormValue("EndDate"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
			return
		}
		notes := r.FormValue("Notes")

		paid, err := strconv.ParseBool(r.FormValue("Paid"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified"))
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
			s.Logger.Errorf("Can't create booking. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat booking with id = %d", b.BookingID)
		http.Redirect(w, r, "/admin/homebookings/", http.StatusFound)

	}
}
