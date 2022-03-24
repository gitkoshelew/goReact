package seathandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/domain/store"
	"hotel/internal/apperror"
	"hotel/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateSeat ...
func CreateSeat(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.SeatDTO{}
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

		seat, err := s.Seat().SeatFromDTO(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("error occured while building model from DTO", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		_, err = s.Seat().Create(seat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while creating seat .", fmt.Sprintf("%d", http.StatusBadRequest),
				fmt.Sprintf("Error occured while creating seat . Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Created seat with id = %d", seat.SeatID)})
	}
}
