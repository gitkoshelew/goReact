package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type seatRequest struct {
	HotelRoomSeatID int    `json:"seatId"`
	Description     string `json:"desc"`
	IsFree          bool   `json:"isFree"`
}

// HandleHotelRoomSeats GET /api/seats - returns all seats(JSON)
//		   	  	  	    POST /api/seat - add seat(JSON)
//				   		PUT /api/seat - update seat(JSON)
func HandleHotelRoomSeats() http.HandlerFunc {

	seats := entity.GetHotelRoomSeats()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		//GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(seats)
		// POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &seatRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			s := entity.HotelRoomSeat{
				HotelRoomSeatID: req.HotelRoomSeatID,
				Description:     req.Description,
				IsFree:          req.IsFree,
			}

			seats = append(seats, s)
			json.NewEncoder(w).Encode(seats)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &seatRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, s := range seats {
				if s.HotelRoomSeatID == req.HotelRoomSeatID {
					seats[index].Description = req.Description
					seats[index].IsFree = req.IsFree
					break
				}
			}
			json.NewEncoder(w).Encode(seats)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleHotelRoomSeat GET /api/seat/:id - returns seat by ID (JSON)
// 		 			   DELETE /api/seat/:id - delete seat by ID(JSON)
func HandleHotelRoomSeat() httprouter.Handle {

	seats := entity.GetHotelRoomSeats()

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
			json.NewEncoder(w).Encode(entity.GetSeatByID(id))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, s := range seats {
				if s.HotelRoomSeatID == id { // delete object imitation =)
					seats[index].Description = "DELETE"
					json.NewEncoder(w).Encode(seats)
					return
				}
			}
			http.Error(w, "Cant find Seat", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
