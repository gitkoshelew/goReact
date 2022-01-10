package pethandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllPetsHandler ...
func AllPetsHandler(s *store.Store) httprouter.Handle {
	//db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		//pets := []model.Pet{}

		s.Open()
		pets, err := s.Pet().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		/*rows, err := db.Query("select * from pet")
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
		*/
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
