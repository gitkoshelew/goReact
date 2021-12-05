package client

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
)

// GetClientsHandler returns all Clients
func GetClientsHandler() http.HandlerFunc {
	clientsDto := dto.GetClientsDto()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(clientsDto)
	}
}
