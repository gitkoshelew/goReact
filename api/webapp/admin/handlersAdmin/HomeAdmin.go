package handlersadmin

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			HomePage(w, s)

			return
		}
		s.Logger.Errorf("Unauthorized")
		http.Redirect(w, r, "/admin/login", http.StatusFound)
	}
}

// HomePage ...
func HomePage(w http.ResponseWriter, s *store.Store) {
	files := []string{
		"/api/webapp/admin/tamplates/homeAdmin.html",
		"/api/webapp/admin/tamplates/base.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), 400)
		s.Logger.Errorf("Can not parse template: %v", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		s.Logger.Errorf("Can not parse template: %v", err)
		return
	}
}
