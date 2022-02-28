package auth

import (
	"admin/domain/store"
	viewdata "admin/pkg/viewData"
	"admin/webapp/session"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		vd := viewdata.ViewData{
			ResponseWriter: w,
			Request:        r,
		}

		files := []string{
			"/rest-api/webapp/tamplates/homeAdmin.html",
			"/rest-api/webapp/tamplates/base.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, vd)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
