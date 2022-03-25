package pethandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permissionCreate model.Permission = model.Permission{Name: model.CreatPet}

// NewPet ...
func NewPet(s *store.Store) httprouter.Handle {
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

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		name := r.FormValue("Name")

		petType := r.FormValue("PetType")

		weight, err := strconv.ParseFloat(r.FormValue("Weight"), 32)

		diseases := r.FormValue("Diseases")

		photo := r.FormValue("Photo")

		petDTO := model.PetDTO{
			PetID:    0,
			Name:     name,
			Type:     petType,
			Weight:   float32(weight),
			Diseases: diseases,
			OwnerID:  userID,
			PhotoURL: photo,
		}

		err = petDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		pet, err := s.Pet().ModelFromDTO(&petDTO)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		_, err = s.Pet().Create(pet)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while creating pet. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/homepets/", http.StatusFound)
	}

}
