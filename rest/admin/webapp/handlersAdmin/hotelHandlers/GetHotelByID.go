package hotelhandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetHotelByID ...
func GetHotelByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_read.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		hotels := []model.Hotel{}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		hotel, err := s.Hotel().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hotels = append(hotels, *hotel)

		files := []string{
			"/rest-api/webapp/tamplates/allHotels.html",
			"/rest-api/webapp/tamplates/base.html",
		}

	tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, hotels)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Can not execute template: %v", err)
			return
		}
	}
}
