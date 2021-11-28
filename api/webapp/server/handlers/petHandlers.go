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

// HandlePets opens a pet page, URL: "/pets". Shows all pets, can search one by id
func HandlePets() http.HandlerFunc {

	pets := webapp.GetPets()

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("webapp/templates/pets.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "pets", pets)
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

		tmpl, err := template.ParseFiles("webapp/templates/show_pet.html", "webapp/templates/header.html", "webapp/templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			log.Fatal(err)
		}
		if petFound {
			tmpl.ExecuteTemplate(w, "show_pet", pet)
		} else {
			tmpl.ExecuteTemplate(w, "show_pet", "Pet not found")
		}

	}
}
