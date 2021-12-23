package hotelhandlers

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"
)

func AllRoomsHandler() http.HandlerFunc {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request) {

		rooms := []store.Room{}

		rows, err := db.Query("select * from room")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			r := store.Room{}
			err := rows.Scan(&r.RoomID, &r.RoomNumber, &r.PetType, &r.Hotel.HotelID)
			if err != nil {
				fmt.Println(err)
				continue
			}
			rooms = append(rooms, r)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allRooms.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, rooms)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}
}
