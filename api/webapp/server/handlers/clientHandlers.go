package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandleClients opens a client page, URL: "/clients". Shows all accounts, can search one by id
func HandleClients() http.HandlerFunc {

	clients := webapp.GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(clients)
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

		if clientFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(client)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
