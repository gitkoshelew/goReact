package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleAccounts opens an account page, URL: "/accounts". Shows all accounts, can search one by id
func HandleAccounts() http.HandlerFunc {

	accounts := webapp.GetAccounts()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
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

		if accountFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(account)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}

	}
}
