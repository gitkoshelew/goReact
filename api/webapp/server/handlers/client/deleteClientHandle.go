package client

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteClientHandle delete Client by ID
func DeleteClientHandle() httprouter.Handle {

	clientsDto := dto.GetClientsDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, c := range clientsDto {
			if c.ClientID == id { // delete object imitation =)
				clientsDto[index].ClientID = 0
				json.NewEncoder(w).Encode(clientsDto)
				return
			}
		}

		http.Error(w, "Cant find Client", http.StatusBadRequest)
	}
}
