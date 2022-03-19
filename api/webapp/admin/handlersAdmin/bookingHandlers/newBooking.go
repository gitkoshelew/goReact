package bookinghandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permissionCreate model.Permission = model.Permission{Name: model.CreatBooking}

// NewBooking ...
func NewBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionCreate.Name)
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
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("SeatID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("SeatID"))
			return
		}
		/*seat, err := s.Seat().FindByID(seatID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			*return
		}*/

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("PetID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("PetID"))
			return
		}
		/*pet, err := s.Pet().FindByID(petID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}*/
		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EmployeeID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EmployeeID"))
			return
		}
		/*employee, err := s.Employee().FindByID(employeeID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting employee by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}*/
		status := r.FormValue("Status")
		layout := "2006-01-02"
		startDate, err := time.Parse(layout, r.FormValue("StartDate"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("StartDate"))
			return
		}
		endDate, err := time.Parse(layout, r.FormValue("EndDate"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EndDate"))
			return
		}
		notes := r.FormValue("Notes")

		paid, err := strconv.ParseBool(r.FormValue("Paid"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Paid"))
			return
		}

		b := model.BookingDTO{
			BookingID: 0,
			SeatID:      seatID,
			PetID:       petID,
			EmployeeID:  employeeID,
			Status:    status,
			StartDate: &startDate,
			EndDate:   &endDate,
			Notes:     notes,
			Paid:      &paid,
		}

		/*err = b.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}*/

		_, err = s.Booking().Create(&b)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while creating booking. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homebookings/", http.StatusFound)

	}
}
