package bookinghandlers

import (
	"fmt"

	"goReact/webapp/admin/session"

	"goReact/domain/model"

	"goReact/webapp/server/utils"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllBookingsHandler ...
func AllBookingsHandler() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
			/*
			s.Open()
		booking, err := s.Booking().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		}*/

		booking := []model.Booking{}
		rows, err := db.Query("select * from booking")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			b := model.Booking{}
			err := rows.Scan(&b.BookingID, &b.Seat.SeatID, &b.Pet.PetID, &b.Employee.EmployeeID, &b.Status, &b.StartDate, &b.EndDate, &b.Notes)
			if err != nil {
				fmt.Println(err)
				continue
			}
			booking = append(booking, b)
		}

		files := []string{
			"/api/webapp/admin/tamplates/allBookings.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, booking)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
