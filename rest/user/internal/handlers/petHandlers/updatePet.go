package pethandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain/model"
	"user/internal/apperror"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// UpdatePet ...
func UpdatePet(s *store.Store) httprouter.Handle {
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
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
		}

		user, err := s.User().FindByID(req.OwnerID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find user.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find user. Err msg:%v.", err)))
			return
		}

		p, err := s.Pet().FindByID(req.PetID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find pet.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find pet. Err msg:%v.", err)))
			return
		}

		if req.Name != "" {
			p.Name = req.Name
		}

		if req.Type != "" {
			p.Type = req.Type
		}

		if req.Weight != 0 {
			p.Weight = req.Weight
		}

		if req.Diseases != "" {
			p.Diseases = req.Diseases
		}

		if user != nil {
			if user.UserID != req.OwnerID {
				p.Owner = *user
			}
		}

		if req.PetPhotoURL != "" {
			p.PetPhotoURL = req.PetPhotoURL
		}

		err = p.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		err = s.Pet().Update(p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't update pet.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Can't update pet. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update pet with id = %d", p.PetID)})
	}
}
