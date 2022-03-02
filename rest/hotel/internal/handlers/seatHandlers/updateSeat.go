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

// UpdateSeat ...
func UpdateSeat(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		roomDTO, err := s.Room().FindByID(req.RoomID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find room.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find room. Err msg:%v.", err)))
			return
		}

		room, err := s.RoomRepository.RoomFromDTO(roomDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't convert room.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't convert room. Err msg:%v.", err)))
			return
		}

		SeatDTO, err := s.Seat().FindByID(req.SeatID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find seat.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find seat. Err msg:%v.", err)))

			return
		}

		seat, err := s.Seat().SeatFromDTO(SeatDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't convert seat.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't convert seat. Err msg:%v.", err)))
			return
		}

		if room != nil {
			if seat.Room.RoomID != req.RoomID {
				seat.Room = *room
			}
		}

		if req.Description != "" {
			seat.Description = req.Description
		}

		if !req.RentFrom.IsZero() {
			seat.RentFrom = req.RentFrom
		}

		if !req.RentTo.IsZero() {
			seat.RentTo = req.RentTo
		}

		err = seat.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		err = s.Seat().Update(seat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't update seat.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Can't update seat. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update seat with id = %d", seat.SeatID)})

	}
}
