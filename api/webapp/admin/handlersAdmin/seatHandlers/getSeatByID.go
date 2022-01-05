package seathandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetSeatByID ...
func GetSeatByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		seats := []model.Seat{}
		id, _ := strconv.Atoi(ps.ByName("id"))

		/*s.Open()
		seats, err := s.Seat().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}*/
		rows, err := db.Query("select * from seat where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			s := model.Seat{}
			err := rows.Scan(&s.SeatID, &s.Room.RoomID, &s.IsFree, &s.Description)
			if err != nil {
				fmt.Println(err)
				continue
			}
			seats = append(seats, s)
		}

		if len(seats) == 0 {
			http.Error(w, "No seat with such id!", 400)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allSeats.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, seats)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
