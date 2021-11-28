package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func HandleAccountsJson() http.HandlerFunc {

	accounts := webapp.GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Sprintf("Json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
	}
}

// HandleAccounts opens an account page, URL: "/accounts". Shows all accounts, can search one by id
func HandleAccounts() http.HandlerFunc {

	accounts := webapp.GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/accounts.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "accounts", accounts)
	}
}

// HandleAccountSearch shows an account by id, URL"/account?id="
func HandleAccountSearch() http.HandlerFunc {

	accounts := webapp.GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var account entity.Account
		accountFound := false

		for _, a := range accounts {
			if a.AccountID == id {
				account = a
				accountFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_account.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if accountFound {
			tmpl.ExecuteTemplate(w, "show_account", account)
		} else {
			tmpl.ExecuteTemplate(w, "show_account", "Account not found")
		}

	}
}
