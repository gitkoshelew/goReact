package auth

import (
	"html/template"
	"net/http"

	"admin/domain/store"
	"admin/webapp/session"

	"github.com/julienschmidt/httprouter"
)

// LoginAdmin ...
func LoginAdmin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		files := []string{
			"/rest-api/webapp/tamplates/loginAdmin.html",
			"/rest-api/webapp/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Can not execute template: %v", err)
			return
		}
	}
}
