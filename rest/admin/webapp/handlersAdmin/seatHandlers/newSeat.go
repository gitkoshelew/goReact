package seathandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permission_create model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "creat_seats",
	Descriptoin:  "ability to creat a seats"}

// NewSeat ...
func NewSeat(s *store.Store) httprouter.Handle {
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
			return
		}

		

		roomID, err := strconv.Atoi(r.FormValue("RoomID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomID"))
			return
		}

		room, err := s.Room().FindByID(roomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't find hotel. Err msg:%v.", err)
			return
		}

		description := r.FormValue("Description")

		layout := "2006-01-02"
		rentFrom, err := time.Parse(layout, r.FormValue("RentFrom"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentFrom"))
			return
		}

		rentTo, err := time.Parse(layout, r.FormValue("RentTo"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentTo"))
			return
		}
		seat := model.Seat{
			SeatID:      0,
			Room:        *room,
			Description: description,
			RentFrom:    rentFrom,
			RentTo:      rentTo,
		}

		err = seat.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Seat().Create(&seat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeseats/", http.StatusFound)

	}

}
