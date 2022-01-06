package employeehandlers

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetEmployeeByID ...
func GetEmployeeByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		employees := []model.Employee{}

		id, _ := strconv.Atoi(ps.ByName("id"))

		s.Open()
		employee, err := s.Employee().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		employees = append(employees, *employee)

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
