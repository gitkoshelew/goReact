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

// HandleClients opens a client page, URL: "/clients". Shows all accounts, can search one by id
func HandleClients() http.HandlerFunc {

	clients := webapp.GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/clients.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "clients", clients)
	}
}

// HandleClientSearch shows a client by id, URL"/client?id="
func HandleClientSearch() http.HandlerFunc {

	clients := webapp.GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var client entity.Client
		clientFound := false

		for _, a := range clients {
			if a.ClientID == id {
				client = a
				clientFound = true
				break
			}
		}

		tmpl, err := template.ParseFiles("webapp/templates/show_client.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if clientFound {
			tmpl.ExecuteTemplate(w, "show_client", client)
		} else {
			tmpl.ExecuteTemplate(w, "show_client", "Client not found")
		}
	}
}
