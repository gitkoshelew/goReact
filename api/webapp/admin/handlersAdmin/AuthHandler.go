package handlersadmin

import (
	"goReact/domain/model"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthAdmin() httprouter.Handle {

	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		EmailForm := r.FormValue("email")
		Password := r.FormValue("password")
		user := model.User{}
		rowPassword := db.QueryRow("SELECT * FROM ACCOUNT WHERE Email = $1 ", EmailForm)
		err := rowPassword.Scan(&user.UserID, &user.Email, &user.Password)

		hashPassword := user.Password
		id := user.UserID

		if err != nil {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return

		}

		isConfirmed := model.CheckPasswordHash(hashPassword, Password)
		if isConfirmed != nil {
			http.Redirect(w, r, "/admin/login", http.StatusFound)

			return
		}
		session.AuthSession(w, r, id)

		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
