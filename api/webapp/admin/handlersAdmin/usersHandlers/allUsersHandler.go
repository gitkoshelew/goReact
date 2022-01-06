package usershandlers

import (
	"goReact/domain/store"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllUsersHandler ...
func AllUsersHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		s.Open()
		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allUsers.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
