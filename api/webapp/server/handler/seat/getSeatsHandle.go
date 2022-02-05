package seat

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllSeatsHandle ...
func GetAllSeatsHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		}

		seats, err := s.Seat().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't find seats. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seats)
	}
}
