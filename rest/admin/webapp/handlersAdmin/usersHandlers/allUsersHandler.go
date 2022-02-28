package usershandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var permission_read model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "read_user",
	Descriptoin:  "ability to read a user"}

// AllUsersHandler ...
func AllUsersHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_read.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}

		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find users. Err msg: %v", err)
			return
		}

		files := []string{
			"/rest-api/webapp/tamplates/allUsers.html",
			"/rest-api/webapp/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
