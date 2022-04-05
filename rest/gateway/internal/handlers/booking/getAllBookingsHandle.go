package booking

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/booking"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllBookingsHandle ...
func GetAllBookingsHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getAllService, err := booking.GetAll(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var booking []*bookingDTO
		if err := json.Unmarshal(getAllService.Body, &booking); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	}
}
