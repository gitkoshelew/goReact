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

func NewSeat(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.SeatDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		s.Logger.Info("req :" , req)
		

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		roomDTO, err := s.Room().FindByID(req.RoomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Cant find room.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Cant find room. Err msg:%v.", err)))
			return
		}
		s.Logger.Info("roomDTO :" , roomDTO)

		room, err := s.Room().RoomFromDTO(roomDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find room.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find room. Err msg:%v.", err)))
			return
		}
		s.Logger.Info("room :" , room)

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
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		_, err = s.Seat().Create(&seat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't create seat.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Can't create seat. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Creat seat with id = %d", seat.SeatID)})
	}
}
