package employeehandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionCreate model.Permission = model.Permission{Name: model.CreatEmployee}

// NewEmployee ...
func NewEmployee(s *store.Store) httprouter.Handle {
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
		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		/*	user, err := s.User().FindByID(userID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error occured while finding user by id. Err msg:%v. ", err), http.StatusBadRequest)
				return
			}*/

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		/*hotel, err := s.Hotel().FindByID(hotelID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while finding hotel by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}*/

		position := r.FormValue("Position")

		e := model.EmployeeDTO{
			EmployeeID: 0,
			UserID:     userID,
			HotelID:    hotelID,
			Position:   position,
		}

		/*err = e.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}*/

		_, err = s.Employee().Create(&e)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while creating employee. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeemployees/", http.StatusFound)
	}

}
