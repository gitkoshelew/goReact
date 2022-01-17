package employeehandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllEmployeeHandler ...
func AllEmployeeHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		s.Open()
		employees, err := s.Employee().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allEmployee.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, employees)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
