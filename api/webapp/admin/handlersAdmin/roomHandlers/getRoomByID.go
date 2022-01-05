package roomhandlers

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

// GetRoomByID ...
func GetRoomByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		rooms := []store.Room{}
		id, _ := strconv.Atoi(ps.ByName("id"))

		rows, err := db.Query("select * from room where id=$1", id)
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

		if len(rooms) == 0 {
			http.Error(w, "No room with such id!", 400)
			return
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
