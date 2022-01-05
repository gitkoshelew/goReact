package bookinghandlers

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

// GetBookingByID ...
func GetBookingByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		booking := []model.Booking{}

		id, _ := strconv.Atoi(ps.ByName("id"))

		/*s.Open()
		pets, err := s.Booking().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}*/
		rows, err := db.Query("select * from booking where id=$1", id)
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
		if len(booking) == 0 {
			http.Error(w, "No booking with such id!", 400)
			return
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
