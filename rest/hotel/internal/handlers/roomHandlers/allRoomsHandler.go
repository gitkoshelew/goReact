package roomhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/store"
	"hotel/internal/apperror"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AllRoomsHandler ...
func AllRoomsHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return

		}
		rooms, err := s.Room().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting all rooms", fmt.Sprintf("%d", http.StatusInternalServerError),
				fmt.Sprintf("Error occured while getting all rooms. Err msg: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rooms)

	}
}
