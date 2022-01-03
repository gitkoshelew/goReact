package hotelhandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func AllHotelsHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		hotels := []store.Hotel{}

		session.CheckSession(w, r)

		rows, err := db.Query("select * from hotel")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			h := store.Hotel{}
			err := rows.Scan(&h.HotelID, &h.Name, &h.Address)
			if err != nil {
				fmt.Println(err)
				continue
			}
			hotels = append(hotels, h)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allHotels.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, hotels)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
