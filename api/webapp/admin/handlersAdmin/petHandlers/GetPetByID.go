package petHandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func GetPetByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		pets := []store.Pet{}

		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from pet where id=$1", id)
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

		if len(pets) == 0 {
			http.Error(w, "No pet with such id!", 400)
			return
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
