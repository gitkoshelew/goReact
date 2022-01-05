package handlersadmin

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/server/utils"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin() httprouter.Handle {

	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var id int
		email := r.FormValue("email")
		Password := r.FormValue("password")
		if email == "" || Password == "" {
			http.Error(w, "Enter email or password", 400)
			return
		}
		user := model.User{}
		rowPassword := db.QueryRow("SELECT * FROM users WHERE email = $1 ", email)
		err := rowPassword.Scan(&user.UserID, &user.Email, &user.Password)

		hashPassword := user.Password
		hashPid := user.UserID
		hashEMail := user.Email

		if err != nil {
			http.Error(w, "Check your email or password1", 400)
			return
		}

		isConfirmed := model.CheckPasswordHash(Password, hashPassword)
		fmt.Println("ac login and pas 3 : ", hashEMail, user.Password, hashPid)
		fmt.Println(isConfirmed)
		// if !isConfirmed {
		// 	http.Error(w, "Check your email or password2", 401)
		// 	return
		// }

		row := db.QueryRow("SELECT * FROM users WHERE email = $1 AND password=$2", email, Password)
		err = row.Scan(&id, &email)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/homeAdmin.html",
			"/api/webapp/admin/tamplates/base.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
