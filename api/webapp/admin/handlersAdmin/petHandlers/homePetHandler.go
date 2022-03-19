package pethandlers

import (
	"goReact/domain/store"
	viewdata "goReact/webapp/admin/pkg/viewData"
	"goReact/webapp/admin/session"
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
			"/api/webapp/admin/tamplates/petHome.html",
			"/api/webapp/admin/tamplates/base.html",
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
