package pethandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_update model.Permission = model.Permission{Name: model.UpdatePet}

// UpdatePet ...
func UpdatePet(s *store.Store) httprouter.Handle {
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

		petID, err := strconv.Atoi(r.FormValue("PetID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		pet, err := s.Pet().FindByID(petID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err == nil {
			if userID != 0 {
				user, err := s.User().FindByID(userID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				pet.Owner = *user
			}
		}

		name := r.FormValue("Name")
		if name != "" {
			pet.Name = name
		}

		petType := r.FormValue("PetType")
		if petType != "" {
			pet.Type = model.PetType(petType)
		}

		weight, err := strconv.ParseFloat(r.FormValue("Weight"), 32)
		if err == nil {
			if weight != 0 {
				pet.Weight = float32(weight)
			}
		}

		diseases := r.FormValue("Diseases")
		if diseases != "" {
			pet.Diseases = diseases
		}

		photo := r.FormValue("Photo")
		if photo != "" {
			pet.PetPhotoURL = photo
		}

		err = pet.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Pet().Update(pet)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homepets/", http.StatusFound)
	}

}
