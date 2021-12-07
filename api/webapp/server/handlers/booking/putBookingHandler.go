package booking

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutBookingsHandler updates Booking
func PutBookingsHandler() http.HandlerFunc {
	bookingsDto := dto.GetBookingsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &bookingRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, b := range bookingsDto {
			if b.BookingID == req.BookingID {
				bookingsDto[index].PetID = req.PetID
				bookingsDto[index].SeatID = req.SeatID
				bookingsDto[index].Status = req.Status
				bookingsDto[index].StartDate = req.StartDate
				bookingsDto[index].EndDate = req.EndDate
				bookingsDto[index].EmployeeID = req.EmployeeID
				bookingsDto[index].ClientNotes = req.ClientNotes
				break
			}
		}
		json.NewEncoder(w).Encode(bookingsDto)
	}
}
