package bookinghandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"

	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllBookingsHandler ...
func AllBookingsHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		booking, err := s.Booking().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find bookings. Err msg: %v", err)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allBookings.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, booking)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
