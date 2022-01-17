package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// HandleHomePage opens a main page, URL: "/"
func HandleHomePage() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tmpl, err := template.ParseFiles("webapp/templates/home_page.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "home_page", nil)
	}
}
