package employeehandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetEmployeeByID ...
func GetEmployeeByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		employees := []model.Employee{}

		id, _ := strconv.Atoi(ps.ByName("id"))

		/*
			s.Open()
			employess, err := s.Employee().FindByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}*/

		rows, err := db.Query("select * from employee where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			e := model.Employee{}
			err := rows.Scan(&e.EmployeeID, &e.User.UserID, &e.Hotel.HotelID, &e.Position, &e.Role)
			if err != nil {
				fmt.Println(err)
				continue
			}
			employees = append(employees, e)
		}

		if len(employees) == 0 {
			http.Error(w, "No employee with such id!", 400)
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
