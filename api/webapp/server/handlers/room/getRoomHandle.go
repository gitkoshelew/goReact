package room

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetRoomHandle returns Room by ID
func GetRoomHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(dto.RoomToDto(entity.GetRoomByID(id)))
	}
}
