package handlersadmin

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthAdmin() httprouter.Handle {

	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		LoginForm := r.FormValue("login")
		Password := r.FormValue("password")
		account := store.Account{}
		rowPassword := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 ", LoginForm)
		err := rowPassword.Scan(&account.AccountID, &account.Login, &account.Password)

		hashPassword := account.Password
		id := account.AccountID
		hashlogin := account.Login
		fmt.Println(Password, hashlogin)

		if err != nil {
			//	http.Error(w, "Check your email or password1", http.StatusBadRequest)
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return

		}

		isConfirmed := store.CheckPasswordHash(hashPassword, Password)
		fmt.Println("hashPassword and pas : ", hashPassword, Password, id)
		fmt.Println("ac login and pas 3 : ", hashlogin, account.Password, id)
		fmt.Println(isConfirmed)
		if isConfirmed != nil {
			//	http.Error(w, "Check your email or password2", http.StatusBadRequest)
			http.Redirect(w, r, "/admin/login", http.StatusFound)

			return
		}
		session.AuthSession(w, r, id)

		http.Redirect(w, r, "/admin/home", http.StatusFound)
		

	}
}
