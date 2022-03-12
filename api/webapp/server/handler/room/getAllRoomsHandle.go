package room

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllRoomsHandle returns all rooms
func GetAllRoomsHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB. Err msg: %v", err)})
			return
		}

		rooms, err := s.Room().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching rooms. Err msg: %v", err)})
			return
		}

		count, err := s.Room().GetTotalRows()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't calculate rows. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		res := make(map[string]interface{})
		res["rooms"] = rooms
		res["totalCount"] = count

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
