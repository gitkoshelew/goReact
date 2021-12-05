package booking

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetBookingHandle returns Booking by ID
func GetBookingHandle() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(entity.BookingToDto(entity.GetBookingByID(id)))
	}
}
