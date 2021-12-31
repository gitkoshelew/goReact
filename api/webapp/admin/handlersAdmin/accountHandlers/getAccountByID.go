package accounthandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetAccountByID ...
func GetAccountByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		accounts := []store.Account{}
		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from account where id=$1", id)
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

		if len(accounts) == 0 {
			http.Error(w, "No account with such id!", 400)
			return
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
