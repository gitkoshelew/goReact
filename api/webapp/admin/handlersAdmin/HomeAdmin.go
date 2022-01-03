package handlersadmin

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func HomeAdmin() httprouter.Handle {

	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//var id int
		LoginForm := r.FormValue("login")
		Password := r.FormValue("password")
		if LoginForm == "" || Password == "" {
			http.Error(w, "Enter email or password", http.StatusBadRequest)
			return
		}
		account := store.Account{}
		rowPassword := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 ", LoginForm)
		err := rowPassword.Scan(&account.AccountID, &account.Login, &account.Password)

		hashPassword := account.Password
		id := account.AccountID
		hashlogin := account.Login
		fmt.Println(Password, hashlogin)

		if err != nil {
			http.Error(w, "Check your email or password1", http.StatusBadRequest)
			return
		}

		isConfirmed := store.CheckPasswordHash(hashPassword, Password)
		fmt.Println("hashPassword and pas : ", hashPassword, Password, id)
		fmt.Println("ac login and pas 3 : ", hashlogin, account.Password, id)
		fmt.Println(isConfirmed)
		if isConfirmed != nil {
			http.Error(w, "Check your email or password2", http.StatusBadRequest)
			return
		}

		session.AuthSession(w, r, id)

		/*session, _ := sessionStore.Get(r, "session")
		session.Values["accountID"] = id
		session.Save(r, w)

		/*row := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 AND password=$2", Login,Password)
		err = row.Scan(&id,&Login)
		if err !=nil{
			http.Error(w, err.Error(), 400)
		return
		}*/

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
