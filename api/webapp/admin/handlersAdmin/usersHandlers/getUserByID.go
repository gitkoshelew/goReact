package usershandlers

import (
	"fmt"

	"goReact/webapp/admin/session"

	"goReact/domain/model"
 
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetUserByID ...
func GetUserByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)

		users := []model.User{}

		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from users where id=$1", id)
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

		if len(users) == 0 {
			http.Error(w, "No user with such id!", 400)
			return
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
