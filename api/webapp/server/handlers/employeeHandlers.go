package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleEmployees opens an employee page, URL: "/employees". Shows all employees, can search one by id
func HandleEmployees() http.HandlerFunc {

	employees := webapp.GetEmployees()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employees)
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

		if employeeFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(employee)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}

	}
}
