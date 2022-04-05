package auth

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginAdmin ...
func LoginAdmin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			http.Redirect(w, r, "/admin/home", http.StatusFound)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/loginAdmin.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}
	}
}
