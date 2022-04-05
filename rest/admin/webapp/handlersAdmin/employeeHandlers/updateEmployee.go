package employeehandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateBooking}

// UpdateEmployee ...
func UpdateEmployee(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionUpdate.Name)
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
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EmployeeID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("EmployeeID"))
			return
		}

		employeeDTO, err := s.Employee().FindByID(employeeID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting employee by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err == nil {
			if userID != 0 {
				employeeDTO.UserID = userID
			}
		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err == nil {
			if hotelID != 0 {
				employeeDTO.HotelID = hotelID
			}
		}

		position := r.FormValue("Position")
		if position != "" {
			employeeDTO.Position = position
		}

		err = employeeDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		employee, err := s.Employee().ModelFromDTO(employeeDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		err = s.Employee().Update(employee)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating employee. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeemployees/", http.StatusFound)
	}

}
