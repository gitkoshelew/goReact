package pethandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain/model"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// NewPet ...
func NewPet(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.PetDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		user, err := s.User().FindByID(req.OwnerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			return
		}

		p := model.Pet{
			PetID:       0,
			Name:        req.Name,
			Type:        model.PetType(req.Type),
			Weight:      req.Weight,
			Diseases:    req.Diseases,
			Owner:       *user,
			PetPhotoURL: req.PetPhotoURL,
		}

		err = p.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		_, err = s.Pet().Create(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create pet. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Creat pet with id = %d", p.PetID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Creat pet with id = %d", p.PetID)})
	}
}
