package pethandlers

import (
	"admin/domain/store"
	viewdata "admin/pkg/viewData"
	"admin/webapp/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// HomePetsHandler ...
func HomePetsHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionRead.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		vd := viewdata.ViewData{
			ResponseWriter: w,
			Request:        r,
		}

		files := []string{
			"/rest-api/webapp/tamplates/petHome.html",
			"/rest-api/webapp/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, vd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}
	}
}
