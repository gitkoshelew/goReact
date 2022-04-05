package handlers

import (
	"booking/domain/model"
	"booking/domain/store"
	"booking/internal/apperror"
	"booking/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateHandle ...
func CreateHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.Booking{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", r.Body), fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		id, err := s.Booking().Create(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while creating booking", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Created booking with id = %d", id)})
	}
}
