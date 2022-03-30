package seathandlers

import (
	"fmt"
	"goReact/domain/request"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

// GetFreeSeatsHandle ...
func GetFreeSeatsHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionRead.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("HotelID"))
			return
		}

		petType := r.FormValue("PetType")

		layout := "2006-01-02"
		rentFrom, err := time.Parse(layout, r.FormValue("RentFrom"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentFrom")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentFrom"))
			return
		}

		rentTo, err := time.Parse(layout, r.FormValue("RentTo"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentTo")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RentTo"))
			return
		}
		req := &request.FreeSeatsSearching{
			HotelID:  hotelID,
			PetType:  petType,
			RentFrom: &rentFrom,
			RentTo:   &rentTo,
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating user. Err msg: %v", err)
			return
		}
		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		seats, err := s.Seat().FreeSeatsSearching(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting all seats. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allSeats.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, seats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while executing template: %v", err)
			return
		}
	}
}
