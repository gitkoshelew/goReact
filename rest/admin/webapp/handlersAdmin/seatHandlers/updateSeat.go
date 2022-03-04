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

var permission_update model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "update_seats",
	Descriptoin:  "ability to update a seats"}

// UpdateSeat ...
func UpdateSeat(s *store.Store) httprouter.Handle {
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

		seatid, err := strconv.Atoi(r.FormValue("SeatID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("SeatID"))
			return
		}

		seat, err := s.Seat().FindByID(seatid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		roomID, err := strconv.Atoi(r.FormValue("RoomID"))
		if err == nil {
			if roomID != 0 {
				room, err := s.Room().FindByID(roomID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					s.Logger.Errorf("Can't find hotel. Err msg:%v.", err)
					return
				}
				seat.Room = *room
			}
		}

		description := r.FormValue("Description")
		if description != "" {
			seat.Description = description
		}

		layout := "2006-01-02"
		rentFrom := r.FormValue("RentFrom")
		if rentFrom != "" {
			rentFrom, err := time.Parse(layout, r.FormValue("RentFrom"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentFrom"))
				return
			}
			seat.RentFrom = rentFrom
		}

		rentTo := r.FormValue("RentTo")
		if rentTo != "" {
			rentTo, err := time.Parse(layout, r.FormValue("RentTo"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentFrom"))
				return
			}
			seat.RentTo = rentTo
		}

		err = seat.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Seat().Update(seat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homeseats/", http.StatusFound)

	}

}
