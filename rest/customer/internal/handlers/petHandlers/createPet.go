package pethandlers

import (
	"customer/domain/model"
	"customer/domain/store"
	"customer/internal/apperror"
	"customer/pkg/response"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreatePet ...
func CreatePet(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.PetDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", r.Body),
				fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		pet, err := s.Pet().ModelFromDTO(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("error occured while building model from DTO", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		id, err := s.Pet().Create(pet)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while creating pet", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Created pet with id = %d", id)})
	}
}
