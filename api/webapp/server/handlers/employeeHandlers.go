package handlers

import (
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// HandleEmployees opens an employee page, URL: "/employees". Shows all employees, can search one by id
func HandleEmployees() http.HandlerFunc {

	employees := webapp.GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/employees.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "employees", employees)
	}
}

// HandleEmployeeSearch shows an employee by id, URL"/employee?id="
func HandleEmployeeSearch() http.HandlerFunc {

	employees := webapp.GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var employee entity.Employee
		employeeFound := false

		for _, a := range employees {
			if a.EmployeeID == id {
				employee = a
				employeeFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_employee.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if employeeFound {
			tmpl.ExecuteTemplate(w, "show_employee", employee)
		} else {
			tmpl.ExecuteTemplate(w, "show_employee", "Employee not found")
		}

	}
}
