package roomhandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionCreate model.Permission = model.Permission{Name: model.CreatRoom}

// NewRoom ...
func NewRoom(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionCreate.Name)
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

		roomNumber, err := strconv.Atoi(r.FormValue("RoomNumber"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomNumber")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomNumber"))
			return
		}
		petType := r.FormValue("PetType")

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		/*hotel, err := s.Hotel().FindByID(hotelID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting hotel by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}*/
		photo := r.FormValue("Photo")

		room := model.RoomDTO{
			RoomID:     0,
			RoomNumber: roomNumber,
			PetType:    petType,
			HotelID:    hotelID,
			PhotoURL:   photo,
		}

		/*err = room.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}*/

		_, err = s.Room().Create(&room)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while creating room. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homerooms/", http.StatusFound)
	}

}
