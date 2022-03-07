package roomhandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_update model.Permission = model.Permission{Name: model.UpdateRoom}

// UpdateRoom ...
func UpdateRoom(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_update.Name)
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

		roomid, err := strconv.Atoi(r.FormValue("RoomID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomID"))
			return
		}

		room, err := s.Room().FindByID(roomid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		roomNumber, err := strconv.Atoi(r.FormValue("RoomNumber"))
		if err == nil {
			if roomNumber != 0 {
				room.RoomNumber = roomNumber
			}
		}

		petType := r.FormValue("PetType")
		if petType != "" {
			room.PetType = model.PetType(petType)

		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err == nil {
			if hotelID != 0 {
				hotel, err := s.Hotel().FindByID(hotelID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				room.Hotel = *hotel
			}
		}
		photo := r.FormValue("Photo")
		if photo != "" {
			room.RoomPhotoURL = photo
		}

		err = room.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Room().Update(room)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homerooms/", http.StatusFound)
	}

}
