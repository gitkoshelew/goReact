package employeehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_update model.Permission = model.Permission{Name: model.UpdateBooking}

// UpdateEmployee ...
func UpdateEmployee(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_update.Name)
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

		employeeID, err := strconv.Atoi(r.FormValue("EmployeeID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		employee, err := s.Employee().FindByID(employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err == nil {
			if userID != 0 {
				user, err := s.User().FindByID(userID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				employee.User = *user
			}
		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err == nil {
			if hotelID != 0 {
				hotel, err := s.Hotel().FindByID(hotelID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				employee.Hotel = *hotel
			}
		}

		position := r.FormValue("Position")
		if position != "" {
			employee.Position = model.Position(position)
		}

		err = employee.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Employee().Update(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeemployees/", http.StatusFound)
	}

}
