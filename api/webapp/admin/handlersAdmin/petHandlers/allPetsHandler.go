package petHandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func AllPetsHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		pets := []store.Pet{}

		rows, err := db.Query("select * from pet")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			p := store.Pet{}
			err := rows.Scan(&p.PetID, &p.Name, &p.Type, &p.Weight, &p.Diesieses, &p.Owner.ClientID)
			if err != nil {
				fmt.Println(err)
				continue
			}
			pets = append(pets, p)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allPets.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, pets)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}
}
