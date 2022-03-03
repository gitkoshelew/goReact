package hotelhandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var permission_create model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "create_hotel",
	Descriptoin:  "ability to create a hotel"}

// NewHotel ...
func NewHotel(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_create.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}
		name := r.FormValue("Name")

		address := r.FormValue("Address")

		coordinates := r.FormValue("Coordinates[0]")

		h := model.Hotel{
			HotelID:     0,
			Name:        name,
			Address:     address,
			Coordinates: coordinates,
		}
		err = h.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Hotel().Create(&h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create hotel. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat hotel with id = %d", h.HotelID)
		http.Redirect(w, r, "/admin/homehotels/", http.StatusFound)
	}

}
