package pethandlers

import (
	"customer/domain/model"
	"customer/internal/apperror"
	"customer/internal/store"
	"customer/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

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
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		user, err := s.User().FindByID(req.OwnerID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting user by id", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		petDTO, err := s.Pet().FindByID(req.PetID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting pet by id", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}
		p, err := s.Pet().PetFromDTO(petDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while convetring petDTO", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
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
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while updating pet.", fmt.Sprintf("%d", http.StatusBadRequest),
				fmt.Sprintf("Can't update pet. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Updated pet with id = %d", p.PetID)})
	}
}
