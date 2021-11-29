package handlers

import (
	"encoding/json"
	"fmt"
	"goReact/domain/entity"
	"goReact/webapp"
	"net/http"
	"strconv"
)

// HandlePets opens a pet page, URL: "/pets". Shows all pets, can search one by id
func HandlePets() http.HandlerFunc {

	pets := webapp.GetPets()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pets)
	}
}

// HandlePetSearch shows a pet by id, URL"/pet?id="
func HandlePetSearch() http.HandlerFunc {

	pets := webapp.GetPets()

	return func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		id, err := strconv.Atoi(urlQuery.Get("id"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		var pet entity.Pet
		petFound := false

		for _, a := range pets {
			if a.PetID == id {
				pet = a
				petFound = true
				break
			}
		}

		if petFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(pet)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}

	}
}
