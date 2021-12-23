package usershandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"
)

func AllUsersHandler() http.HandlerFunc {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request) {

		users := []store.User{}

		rows, err := db.Query("select * from users")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			u := store.User{}
			err := rows.Scan(&u.UserID, &u.Name, &u.Surname, &u.MiddleName, &u.Email, &u.DateOfBirth, &u.Address, &u.Phone, &u.Account.AccountID)
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
