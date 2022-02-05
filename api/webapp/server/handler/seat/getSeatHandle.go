package seat

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetSeatHandle ...
func GetSeatHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		seat, err := s.Seat().FindByID(id)

		s.Logger.Debugf("Searching seat by id: %d", id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant find seat. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		s.Logger.Debugf("Seat: %v", seat)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seat)
	}
}
