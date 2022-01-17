package bookinghandlers

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetBookingByID ...
func GetBookingByID(s *store.Store) httprouter.Handle { //db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		bookings := []model.Booking{}

		id, _ := strconv.Atoi(ps.ByName("id"))

		s.Open()
		booking, err := s.Booking().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		bookings = append(bookings, *booking)

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
