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

		user, err := s.User().FindByID(req.OwnerID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting pet by id", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		p := model.Pet{
			PetID:    0,
			Name:     req.Name,
			Type:     model.PetType(req.Type),
			Weight:   req.Weight,
			Diseases: req.Diseases,
			Owner:    *user,
			PhotoURL: req.PhotoURL,
		}

		_, err = s.Pet().Create(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while creating pet", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Created pet with id = %d", p.PetID)})
	}
}
