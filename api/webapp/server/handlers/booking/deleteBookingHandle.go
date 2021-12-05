package booking

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteBookingHandle returns Booking by ID
func DeleteBookingHandle() httprouter.Handle {
	bookingsDto := dto.GetBookingsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, b := range bookingsDto {
			if b.BookingID == id { // delete object imitation =)
				bookingsDto[index].ClientNotes = "DELETE"
				json.NewEncoder(w).Encode(bookingsDto)
				return
			}
		}
		http.Error(w, "Cant find Booking", http.StatusBadRequest)
	}
}
