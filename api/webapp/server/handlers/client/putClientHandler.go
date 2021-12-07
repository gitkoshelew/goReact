package client

import (
	"encoding/json"
	"goReact/domain/entity"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PutClientsHandler updates Client
func PutClientsHandler() http.HandlerFunc {

	clientsDto := dto.GetClientsDto()

	return func(w http.ResponseWriter, r *http.Request) {
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

		clientsDto = append(clientsDto, dto.ClientToDto(c))
		json.NewEncoder(w).Encode(clientsDto)
	}
}
