package image

import (
	"encoding/json"
	"fmt"
	"image/domain/store"
	"image/internal/apperror"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllImagesHandle ...
func GetAllImagesHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		imgs, err := s.Image().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting all images", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while getting hotel by id. Err msg: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(imgs)
	}
}
