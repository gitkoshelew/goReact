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

	seatsDto := entity.GetHotelRoomSeatsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		//GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(seatsDto)
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

			seatsDto = append(seatsDto, entity.SeatToDto(s))
			json.NewEncoder(w).Encode(seatsDto)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &seatRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, s := range seatsDto {
				if s.HotelRoomSeatID == req.HotelRoomSeatID {
					seatsDto[index].Description = req.Description
					seatsDto[index].IsFree = req.IsFree
					break
				}
			}
			json.NewEncoder(w).Encode(seatsDto)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleHotelRoomSeat GET /api/seat/:id - returns seat by ID (JSON)
// 		 			   DELETE /api/seat/:id - delete seat by ID(JSON)
func HandleHotelRoomSeat() httprouter.Handle {

	seats := entity.GetHotelRoomSeatsDto()

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
			json.NewEncoder(w).Encode(entity.SeatToDto(entity.GetSeatByID(id)))
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
