package handlersadmin

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
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
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		id := user.UserID
		hashPassword := user.Password

		isConfirmed := model.CheckPasswordHash(hashPassword, Password)
		if isConfirmed != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		session.AuthSession(w, r, id)

		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
