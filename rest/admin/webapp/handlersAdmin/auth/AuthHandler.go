package auth

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
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
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Error(w, "Eror during checking users email or password", 400)
			return
		}

		userID := user.UserID
		hashPassword := user.Password

		err = model.CheckPasswordHash(hashPassword, Password)
		if err != nil {
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Error(w, "Eror during checking users email or password", 400)
			return
		}

		if user.Role != "employee" {
			http.Error(w, "You are not employee", http.StatusForbidden)
			s.Logger.Errorf("You are not employee")
			return
		}

		employee, err := s.Employee().FindByUserID(userID)
		if err != nil {
			http.Error(w, "No such employee", http.StatusBadRequest)
			s.Logger.Errorf("Eror during getting employee. Err msg: %s", err.Error())
			return
		}

		permissions, err := s.Permissions().GetByEmployeeId(employee.EmployeeID)
		if err != nil {
			http.Error(w, "No such permossions", http.StatusBadRequest)
			s.Logger.Errorf("Eror during getting permissions. Err msg: %s", err.Error())
			return
		}

		session.AuthSession(w, r, employee, permissions)

		http.Redirect(w, r, "/", http.StatusFound)

	}
}
