package seathandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/admin/session"
	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllSeatsHandler ...
func AllSeatsHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		seats := []model.Seat{}

		rows, err := db.Query("select * from seat")
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
