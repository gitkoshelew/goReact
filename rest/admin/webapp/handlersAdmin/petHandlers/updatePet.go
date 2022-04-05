package pethandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionUpdate model.Permission = model.Permission{Name: model.UpdatePet}

// UpdatePet ...
func UpdatePet(s *store.Store) httprouter.Handle {
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

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("PetID"))
			return
		}

		petDTO, err := s.Pet().FindByID(petID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while getting pet by id. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err == nil {
			if userID != 0 {

				petDTO.OwnerID = userID
			}
		}

		name := r.FormValue("Name")
		if name != "" {
			petDTO.Name = name
		}

		petType := r.FormValue("PetType")
		if petType != "" {
			petDTO.Type = petType
		}

		weight, err := strconv.ParseFloat(r.FormValue("Weight"), 32)
		if err == nil {
			if weight != 0 {
				petDTO.Weight = float32(weight)
			}
		}

		diseases := r.FormValue("Diseases")
		if diseases != "" {
			petDTO.Diseases = diseases
		}

		photo := r.FormValue("Photo")
		if photo != "" {
			petDTO.PhotoURL = photo
		}

		err = petDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}
		pet, err := s.Pet().ModelFromDTO(petDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		err = s.Pet().Update(pet)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating pet. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homepets/", http.StatusFound)
	}

}
