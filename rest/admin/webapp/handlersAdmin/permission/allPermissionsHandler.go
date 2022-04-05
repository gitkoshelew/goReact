package permission

import (
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllPermissons ...
func AllPermissons(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, "admin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf(" Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		per, err := s.Permissions().GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting all permissions. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allPermissions.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, per)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}

	}
}
