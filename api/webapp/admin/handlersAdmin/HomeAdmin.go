package handlersadmin

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"html/template"
	"net/http"
)

func HomeAdmin() http.HandlerFunc {

	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		//var password1 string
		Login := r.FormValue("login")
		Password := r.FormValue("password")
		if Login == "" || Password == ""  {
			http.Error(w, "Enter email or password", 400)
			return
		}
		fmt.Println("ac login and pas 1 : ", Login, Password)
		account := store.Account{}
		rowPassword := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 ", Login)
		err:= rowPassword.Scan( &account.AccountID,&account.Login, &account.Password )

		fmt.Println("ac login and pas 2 : " ,account.Login, account.Password)

		hashPassword := account.Password
		hashPid := account.AccountID
		hashlogin := account.Login
		
		
		if err !=nil{
			http.Error(w, "Check your email or password1", 400)
			return
		}

		isConfirmed := store.CheckPasswordHash(Password , hashPassword)
		fmt.Println("ac login and pas 3 : " ,hashlogin, account.Password, hashPid)
		fmt.Println(isConfirmed)
		if !isConfirmed {
			http.Error(w, "Check your email or password2", 401)
			return
		}

		/*row := db.QueryRow("SELECT * FROM ACCOUNT WHERE login = $1 AND password=$2", Login,Password)
		err = row.Scan(&id,&Login)
		if err !=nil{
			http.Error(w, err.Error(), 400)	
		return
		}*/


		

		tmpl, err := template.ParseFiles("/api/webapp/admin/tamplates/homeAdmin.html")
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
