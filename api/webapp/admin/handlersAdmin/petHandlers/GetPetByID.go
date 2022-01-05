package pethandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetPetByID ...
func GetPetByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		pets := []model.Pet{}

		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from pet where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			p := model.Pet{}
			err := rows.Scan(&p.PetID, &p.Name, &p.Type, &p.Weight, &p.Diesieses, &p.Owner.UserID)
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
