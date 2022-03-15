package auth

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthAdmin ...
func AuthAdmin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		EmailForm := r.FormValue("email")
		Password := r.FormValue("password")
		s.Open()
		user, err := s.User().FindByEmail(EmailForm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while checking users email or password. Err msg:%v. ", err), http.StatusBadRequest)
			s.Logger.Errorf("Error occured while checking users email or password. Err msg: %s", err.Error())
			return
		}

		userID := user.UserID
		hashPassword := user.Password

		err = model.CheckPasswordHash(hashPassword, Password)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while checking users email or password. Err msg:%v. ", err), http.StatusBadRequest)
			s.Logger.Errorf("Error occured while checking users email or password. Err msg: %s", err.Error())
			return
		}

		if user.Role != "employee" {
			http.Error(w, "Access is denied", http.StatusForbidden)
			s.Logger.Errorf("Access is denied")
			return
		}

		employee, err := s.Employee().FindByUserID(userID)
		if err != nil {
			http.Error(w, "Error occured while getting employee", http.StatusInternalServerError)
			return
		}

		permissions, err := s.Permissions().GetEmployeeByID(employee.EmployeeID)
		if err != nil {
			http.Error(w, "Error occured while getting permossions", http.StatusInternalServerError)
			return
		}

		session.AuthSession(w, r, employee, permissions)

		http.Redirect(w, r, "/", http.StatusFound)

	}
}
