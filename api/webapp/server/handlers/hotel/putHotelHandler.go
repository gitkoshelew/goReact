package hotel

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutHotelHandler updates Hotel
func PutHotelHandler() http.HandlerFunc {
	hotels := dto.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &hotelRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, h := range hotels {
			if h.HotelID == req.HotelID {
				hotels[index].Name = req.Name
				hotels[index].Address = req.Address
				hotels[index].RoomsID = req.RoomsID
				hotels[index].BookingsID = req.BookingsID
				break
			}
		}

		json.NewEncoder(w).Encode(hotels)
	}
}
