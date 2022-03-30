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

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateRoom}

// UpdateRoom ...
func UpdateRoom(s *store.Store) httprouter.Handle {
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

		roomID, err := strconv.Atoi(r.FormValue("RoomID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("RoomID"))
			return
		}

		roomDTO, err := s.Room().FindByID(roomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		roomNumber, err := strconv.Atoi(r.FormValue("RoomNumber"))
		if err == nil {
			if roomNumber != 0 {
				roomDTO.RoomNumber = roomNumber
			}
		}

		petType := r.FormValue("PetType")
		if petType != "" {
			roomDTO.PetType = petType

		}

		hotelID, err := strconv.Atoi(r.FormValue("HotelID"))
		if err == nil {
			if hotelID != 0 {
				roomDTO.HotelID = hotelID
			}
		}
		photo := r.FormValue("Photo")
		if photo != "" {
			roomDTO.PhotoURL = photo
		}

		description := r.FormValue("Description")
		if description != "" {
			roomDTO.Description = description
		}

		square, err := strconv.ParseFloat(r.FormValue("Square"), 32)
		if square != 0 {
			roomDTO.Square = square
		}

		err = roomDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}
		room, err := s.Room().ModelFromDTO(roomDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		err = s.Room().Update(room)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating room. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homerooms/", http.StatusFound)
	}

}
