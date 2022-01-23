package usershandlers

import (
	"goReact/webapp/admin/session"

	"goReact/domain/model"
	"goReact/domain/store"

	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetUserByID ...
func GetUserByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)

		users := []model.User{}

		id, _ := strconv.Atoi(r.FormValue("id"))

		s.Open()
		user, err := s.User().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		users = append(users, *user)

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
