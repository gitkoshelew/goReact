package handlersadmin

import (
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

		if err != nil {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return

		}

		isConfirmed := store.CheckPasswordHash(hashPassword, Password)
		if isConfirmed != nil {
			http.Redirect(w, r, "/admin/login", http.StatusFound)

			return
		}
		session.AuthSession(w, r, id)

		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
