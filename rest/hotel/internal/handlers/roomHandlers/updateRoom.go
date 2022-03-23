package roomhandlers

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

// UpdateRoom ...
func UpdateRoom(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.RoomDTO{}
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

		roomDTO, err := s.Room().FindByID(req.RoomID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting room by id.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Error occured while getting room by id. Err msg:%v.", err)))
			return
		}

		room, err := s.RoomRepository.ModelFromDTO(roomDTO)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't convert room.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't convert room. Err msg:%v.", err)))
			return
		}

		err = s.Room().Update(room)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't update room.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Can't update room. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Updated room with id = %d", room.RoomID)})

	}
}
