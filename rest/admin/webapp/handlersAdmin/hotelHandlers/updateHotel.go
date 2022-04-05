package hotelhandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateHotel}

// UpdateHotel ...
func UpdateHotel(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionUpdate.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		hotel, err := s.Hotel().FindByID(hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name := r.FormValue("Name")
		if name != "" {
			hotel.Name = name
		}

		address := r.FormValue("Address")
		if address != "" {
			hotel.Address = address
		}

		lat, err := strconv.ParseFloat(r.FormValue("Lat"), 32)
		if err == nil {
			if lat != 0 {
				hotel.Coordinates[0] = lat
			}
		}

		lon, err := strconv.ParseFloat(r.FormValue("Lon"), 32)
		if err == nil {
			if lon != 0 {
				hotel.Coordinates[1] = lon
			}
		}

		err = hotel.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Hotel().Update(hotel)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating hotel. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/homehotels/", http.StatusFound)
	}

}
