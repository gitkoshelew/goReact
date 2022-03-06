package hotel

import (
	"context"
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/hotel"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetRoomHandle ...
func GetRoomHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getService, err := hotel.Get(context.WithValue(r.Context(), client.HotelGetQuerryParamsCtxKey, ps.ByName("id")), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var room *RoomDTO
		if err := json.Unmarshal(getService.Body, &room); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(room)
	}
}
