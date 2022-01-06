package handlersadmin

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		EmailForm := r.FormValue("email")
		Password := r.FormValue("password")
		s.Open()
		user, _ := s.User().FindByEmail(EmailForm)
		id := user.UserID

		hashPassword := user.Password

		isConfirmed := model.CheckPasswordHash(hashPassword, Password)
		if isConfirmed != nil {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		session.AuthSession(w, r, id)

		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
