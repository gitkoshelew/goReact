package permission

import (
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// HomePetsHandler ...
func HomePermissions(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, "admin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf(" Err msg:%v. ", err)
			return
		}

		files := []string{
			"/rest-api/webapp/tamplates/permissionsHome.html",
			"/rest-api/webapp/tamplates/base.html",
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
}
