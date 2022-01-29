package room

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetRoomHandle returns Room by ID
func GetRoomHandle(s *store.Store) httprouter.Handle {
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

		room, err := s.Room().FindByID(id)

		s.Logger.Debugf("Searching room by id: %d", id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant find room. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		s.Logger.Debugf("Room: %v", room)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(room)
	}
}
