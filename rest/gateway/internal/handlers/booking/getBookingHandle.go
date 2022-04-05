package booking

import (
	"context"
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/booking"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetBookingHandle ...
func GetBookingHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getService, err := booking.Get(context.WithValue(r.Context(), client.BookingGetQuerryParamsCtxKey, ps.ByName("id")), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var booking *bookingDTO
		if err := json.Unmarshal(getService.Body, &booking); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	}
}
