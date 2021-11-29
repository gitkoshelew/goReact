package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleUsers opens an user page, URL: "/users". Shows all user, can search one by id
func HandleUsers() http.HandlerFunc {

	users := webapp.GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

// HandleUserSearch shows an user by id, URL"/user?id="
func HandleUserSearch() http.HandlerFunc {

	users := webapp.GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var user entity.User
		userFound := false

		for _, a := range users {
			if a.UserID == id {
				user = a
				userFound = true
				break
			}
		}

		if userFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}

	}
}
