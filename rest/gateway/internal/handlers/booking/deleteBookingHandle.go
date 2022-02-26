package booking

import (
	"context"
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/booking"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DeleteBookingHandle ...
func DeleteBookingHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		deleteService, err := booking.Delete(context.WithValue(r.Context(), client.BookingDeleteQuerryParamsCtxKey, ps.ByName("id")), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var response *response.Info
		if err := json.Unmarshal(deleteService.Body, &response); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
