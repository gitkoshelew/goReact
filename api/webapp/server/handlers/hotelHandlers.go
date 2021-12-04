package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type hotelRequest struct {
	HotelID    int    `json:"hotelId"`
	Name       string `json:"nameId"`
	Address    string `json:"addressId"`
	RoomsID    []int  `json:"roomsIds"`
	BookingsID []int  `json:"bookingsIds"`
}

// HandleHotels  GET /api/hotels - returns all hotels(JSON)
//	   	  	     POST /api/hotel - add hotel(JSON)
//			     PUT /api/hotel - update hotel(JSON)
func HandleHotels() http.HandlerFunc {

	hotels := entity.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(hotels)
		// POST
		case http.MethodPost:
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
			hotels = append(hotels, entity.HotelToDto(h))
			json.NewEncoder(w).Encode(hotels)
		// PUT
		case http.MethodPut:
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
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleHotel GET /api/hotel/:id - returns hotel by ID (JSON)
// 		 	   DELETE /api/hotel/:id - delete hotel by ID(JSON)
func HandleHotel() httprouter.Handle {

	hotels := entity.GetHotelsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
			json.NewEncoder(w).Encode(entity.HotelToDto(entity.GetHotelByID(id)))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, h := range hotels {
				if h.HotelID == id { // delete object imitation =)
					hotels[index].Name = "DELETE"
					json.NewEncoder(w).Encode(hotels)
					return
				}
			}
			http.Error(w, "Cant find Hotel", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
