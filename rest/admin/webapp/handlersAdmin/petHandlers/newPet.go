package pethandlers

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
	Name:         "creat_pet",
	Descriptoin:  "ability to create a pet"}

// NewPet ...
func NewPet(s *store.Store) httprouter.Handle {
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

		userID, err := strconv.Atoi(r.FormValue("UserID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("UserID"))
			return
		}

		user, err := s.User().FindByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find hotel. Err msg:%v.", err)
			return
		}

		name := r.FormValue("Name")

		petType := r.FormValue("PetType")

		weight, err := strconv.ParseFloat(r.FormValue("Weight"), 2)

		diseases := r.FormValue("Diseases")

		photo := r.FormValue("Photo")

		pet := model.Pet{
			PetID:       0,
			Name:        name,
			Type:        model.PetType(petType),
			Weight:      float32(weight),
			Diseases:    diseases,
			Owner:       *user,
			PetPhotoURL: photo,
		}

		err = pet.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Pet().Create(&pet)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create pet. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat pet with id = %d", pet.PetID)
		http.Redirect(w, r, "/admin/homepets/", http.StatusFound)
	}

}
