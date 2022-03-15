package customer

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/customer"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllPetsHandle ...
func GetAllPetsHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getAllService, err := customer.GetAll(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var pets []*PetDTO
		if err := json.Unmarshal(getAllService.Body, &pets); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pets)
	}
}
