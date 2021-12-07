package hotel

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostHotelHandler creates Hotel
func PostHotelHandler() http.HandlerFunc {
	hotels := dto.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &hotelRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		h := entity.Hotel{
			HotelID:  req.HotelID,
			Name:     req.Name,
			Address:  req.Address,
			Rooms:    entity.GetRoomsByID(req.RoomsID),
			Bookings: entity.GetBookingsByID(req.BookingsID),
		}

		hotels = append(hotels, dto.HotelToDto(h))
		json.NewEncoder(w).Encode(hotels)
	}
}
