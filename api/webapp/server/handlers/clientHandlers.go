package handlers

import (
	"encoding/json"
	"goReact/domain/entity"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type clientRequest struct {
	UserID      int   `json:"userId"`
	ClientID    int   `json:"clientId"`
	PetsIDs     []int `json:"petIds"`
	BookingsIDs []int `json:"bookingIds"`
}

// HandleClients  GET /api/clients - returns all clients(JSON)
//	   	  	      POST /api/client - add client(JSON)
//			 	  PUT /api/client - update client(JSON)
func HandleClients() http.HandlerFunc {

	clients := entity.GetClients()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		//GET
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(clients)
		//POST
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &clientRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			c := entity.Client{
				User:     entity.GetUserByID(req.UserID),
				ClientID: req.ClientID,
				Pets:     entity.GetPetsByID(req.PetsIDs),
				Bookings: entity.GetBookingsByID(req.BookingsIDs),
			}
			clients = append(clients, c)
			json.NewEncoder(w).Encode(clients)
		// PUT
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			req := &clientRequest{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, c := range clients {
				if c.ClientID == req.ClientID {
					clients[index].User = entity.GetUserByID(req.UserID)
					clients[index].Pets = entity.GetPetsByID(req.PetsIDs)
					clients[index].Bookings = entity.GetBookingsByID(req.BookingsIDs)
					break
				}
			}
			json.NewEncoder(w).Encode(clients)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// HandleClient GET /api/client/:id - returns client by ID (JSON)
// 				DELETE /api/client/:id - delete client by ID(JSON)
func HandleClient() httprouter.Handle {

	clients := entity.GetClients()

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

			json.NewEncoder(w).Encode(entity.GetClientByID(id))
		// DELETE
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			id, err := strconv.Atoi(ps.ByName("id"))

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			for index, c := range clients {
				if c.ClientID == id { // delete object imitation =)
					clients[index].ClientID = 0
					json.NewEncoder(w).Encode(clients)
					return
				}
			}
			http.Error(w, "Cant find Client", http.StatusBadRequest)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
