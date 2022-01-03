package handlersadmin

import (
	"fmt"
	"goReact/domain/store"
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
		Login := r.FormValue("login")
		Password := r.FormValue("password")
		if Login == "" || Password == "" {
			http.Error(w, "Enter email or password", 400)
			return
		}
		account := store.Account{}
		rowPassword := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 ", Login)
		err := rowPassword.Scan(&account.AccountID, &account.Login, &account.Password)

		hashPassword := account.Password
		hashPid := account.AccountID
		hashlogin := account.Login

		if err != nil {
			http.Error(w, "Check your email or password1", 400)
			return
		}

		isConfirmed := store.CheckPasswordHash(Password, hashPassword)
		fmt.Println("ac login and pas 3 : ", hashlogin, account.Password, hashPid)
		fmt.Println(isConfirmed)
		// if !isConfirmed {
		// 	http.Error(w, "Check your email or password2", 401)
		// 	return
		// }

		row := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 AND password=$2", Login, Password)
		err = row.Scan(&id, &Login)
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
