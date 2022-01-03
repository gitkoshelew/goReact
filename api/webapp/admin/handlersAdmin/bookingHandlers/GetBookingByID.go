package bookinghandlers

import (
	"fmt"
	"goReact/domain/store"
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

		bookings := []store.Booking{}

		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from booking where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			b := store.Booking{}
			err := rows.Scan(&b.BookingID, &b.Seat.SeatID, &b.Pet.PetID, &b.Employee.EmployeeID, &b.Status, &b.StartDate, &b.EndDate, &b.ClientNotes)
			if err != nil {
				fmt.Println(err)
				continue
			}
			bookings = append(bookings, b)

		}
		if len(bookings) == 0 {
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

		err = tmpl.Execute(w, bookings)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
