package hotelhandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// AllHotelsHandler ...
func AllHotelsHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		s.Open()
		hotels, err := s.Hotel().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
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
