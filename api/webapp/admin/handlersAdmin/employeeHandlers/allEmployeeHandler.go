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

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}

		employees, err := s.Employee().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)			
			s.Logger.Errorf("Can't find employess. Err msg: %v", err)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allEmployee.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, employees)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
