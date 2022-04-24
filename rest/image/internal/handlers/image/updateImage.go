package image

import (
	"encoding/json"
	"fmt"
	"image/domain/model"
	"image/domain/store"
	"image/internal/apperror"
	"image/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UpdateImage ...
func UpdateImage(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.ImageDTO{}
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

		err = req.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		image, err := s.Image().ModelFromDTO(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while converting DTO. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		err = s.Image().Update(image)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating image. Err msg:%v. ", err), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Updated image with id = %d", req.ImageID)})

	}
}
