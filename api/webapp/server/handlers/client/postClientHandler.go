package client

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// PostClientsHandler creates Client
func PostClientsHandler() http.HandlerFunc {

	clientsDto := dto.GetClientsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		req := &clientRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, c := range clientsDto {
			if c.ClientID == req.ClientID {
				clientsDto[index].PetsID = req.PetsIDs
				clientsDto[index].BookingsID = req.BookingsIDs
				break
			}
		}

		json.NewEncoder(w).Encode(clientsDto)
	}
}
