package accounthandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AllAccountsHandler ...
func AllAccountsHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		accounts := []store.Account{}

		rows, err := db.Query("select * from account")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			a := store.Account{}
			err := rows.Scan(&a.AccountID, &a.Login, &a.Password)
			if err != nil {
				fmt.Println(err)
				continue
			}
			accounts = append(accounts, a)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allAccounts.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, accounts)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
