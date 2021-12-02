package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type roomRequest struct {
	HotelRoomID int    `json:"roomId"`
	RoomNumber  int    `json:"roomNum"`
	PetType     string `json:"petType"`
	Seats       []int  `json:"seatsIds"`
}

// HandleHotelRooms GET /api/rooms - returns all rooms(JSON)
//	   	  	  	    POST /api/room - add room(JSON)
//			   		PUT /api/room - update room(JSON)
func HandleHotelRooms() http.HandlerFunc {

	rooms := entity.GetHotelRooms()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rooms)
		// POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &roomRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			r := entity.HotelRoom{
				HotelRoomID: req.HotelRoomID,
				RoomNumber:  req.RoomNumber,
				PetType:     entity.PetType(req.PetType),
				Seats:       entity.GetSeatsByID(req.Seats),
			}
			rooms = append(rooms, r)
			json.NewEncoder(w).Encode(rooms)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &roomRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, r := range rooms {
				if r.HotelRoomID == req.HotelRoomID {
					rooms[index].PetType = entity.PetType(req.PetType)
					rooms[index].RoomNumber = req.RoomNumber
					rooms[index].Seats = entity.GetSeatsByID(req.Seats)
					break
				}
			}
			json.NewEncoder(w).Encode(rooms)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleHotelRoom GET /api/room/:id - returns room by ID (JSON)
// 		 	   DELETE /api/room/:id - delete room by ID(JSON)
func HandleHotelRoom() httprouter.Handle {

	rooms := entity.GetHotelRooms()

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
			json.NewEncoder(w).Encode(entity.GetRoomByID(id))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, r := range rooms {
				if r.HotelRoomID == id { // delete object imitation =)
					rooms[index].HotelRoomID = 0
					json.NewEncoder(w).Encode(rooms)
					return
				}
			}
			http.Error(w, "Cant find Room", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
