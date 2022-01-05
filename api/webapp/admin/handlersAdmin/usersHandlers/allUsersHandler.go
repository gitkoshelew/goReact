package usershandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllUsersHandler ...
func AllUsersHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		users := []model.User{}

		/*s.Open()
		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}*/

		rows, err := db.Query("select * from users")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			u := model.User{}
			err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.Role, &u.Verified, &u.Name, &u.Surname, &u.MiddleName, &u.Sex, &u.DateOfBirth,
				&u.Address, &u.Phone, &u.Photo)

			if err != nil {
				fmt.Println(err)
				continue
			}
			users = append(users, u)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allUsers.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
