package employeehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_create model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "create_employees",
	Descriptoin:  "ability to create a employees"}

// NewEmployee ...
func NewEmployee(s *store.Store) httprouter.Handle {
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
		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		user, err := s.User().FindByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find hotel. Err msg:%v.", err)
			return
		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		hotel, err := s.Hotel().FindByID(hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find hotel. Err msg:%v.", err)
			return
		}

		position := r.FormValue("Position")

		e := model.Employee{
			EmployeeID: 0,
			User:       *user,
			Hotel:      *hotel,
			Position:   model.Position(position),
		}

		err = e.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Employee().Create(&e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create employee. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat employee with id = %d", e.EmployeeID)
		http.Redirect(w, r, "/admin/homeemployees/", http.StatusFound)
	}

}
