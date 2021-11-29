package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// HandleHomePage opens a main page, URL: "/"
func HandleHomePage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/home_page.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "home_page", nil)
	}
}
