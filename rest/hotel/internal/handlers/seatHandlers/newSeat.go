package seathandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/apperror"
	"hotel/internal/store"
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

		s.Logger.Info("req :", req)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		roomDTO, err := s.Room().FindByID(req.RoomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting seat by id.", fmt.Sprintf("%d", http.StatusBadRequest),
				fmt.Sprintf("Error occured while getting seat by id. Err msg:%v.", err)))
			return
		}
		s.Logger.Info("roomDTO :", roomDTO)

		room, err := s.Room().RoomFromDTO(roomDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while convetring roomDTO.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while convetring roomDTO. Err msg:%v.", err)))
			return
		}
		s.Logger.Info("room :", room)

		seat := model.Seat{
			SeatID:      0,
			Room:        *room,
			Description: req.Description,
			RentFrom:    req.RentFrom,
			RentTo:      req.RentTo,
		}

		err = seat.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest),
				fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		_, err = s.Seat().Create(&seat)
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
