package booking

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostBookingsHandler creates Booking
func PostBookingsHandler() http.HandlerFunc {
	bookingsDto := dto.GetBookingsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &bookingRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		b := entity.Booking{
			BookingID:   req.BookingID,
			Pet:         entity.GetPetByID(req.PetID),
			Seat:        entity.GetSeatByID(req.SeatID),
			Status:      entity.BookingStatus(req.Status),
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
			Employee:    entity.GetEmployeeByID(req.EmployeeID),
			ClientNotes: req.ClientNotes,
		}
		bookingsDto = append(bookingsDto, dto.BookingDto(entity.BookingToDto(b)))
		json.NewEncoder(w).Encode(bookingsDto)
	}
}
