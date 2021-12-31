package employeehandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllEmployeeHandler ...
func AllEmployeeHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		employees := []store.Employee{}

		rows, err := db.Query("select * from employee")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			e := store.Employee{}
			err := rows.Scan(&e.EmployeeID, &e.User.UserID, &e.Hotel.HotelID, &e.Position, &e.Role)
			if err != nil {
				fmt.Println(err)
				continue
			}
			employees = append(employees, e)
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
