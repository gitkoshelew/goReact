package roomhandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_create model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "creat_rooms",
	Descriptoin:  "ability to create a rooms"}

// NewRoom ...
func NewRoom(s *store.Store) httprouter.Handle {
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

		roomNumber, err := strconv.Atoi(r.FormValue("RoomNumber"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomNumber"))
			return
		}
		petType := r.FormValue("PetType")

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		hotel, err := s.Hotel().FindByID(hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find hotel. Err msg:%v.", err)
			return
		}
		photo := r.FormValue("Photo")

		room := model.Room{
			RoomID:       0,
			RoomNumber:   roomNumber,
			PetType:      model.PetType(petType),
			Hotel:        *hotel,
			RoomPhotoURL: photo,
		}

		err = room.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Room().Create(&room)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create room. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat room with id = %d", room.RoomID)
		http.Redirect(w, r, "/admin/homerooms/", http.StatusFound)
	}

}
