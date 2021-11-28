package handlers

import (
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// HandleUsers opens an user page, URL: "/users". Shows all user, can search one by id
func HandleUsers() http.HandlerFunc {

	users := webapp.GetUsers()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/users.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "users", users)
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

		tmpl, err := template.ParseFiles("webapp/templates/show_user.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if userFound {
			tmpl.ExecuteTemplate(w, "show_user", user)
		} else {
			tmpl.ExecuteTemplate(w, "show_user", "User not found")
		}

	}
}
