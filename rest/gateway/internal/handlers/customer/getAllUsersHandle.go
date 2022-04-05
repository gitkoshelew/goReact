package customer

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/customer"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllUsersHandle ...
func GetAllUsersHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getAllService, err := customer.GetAll(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var users []*UserDTO
		if err := json.Unmarshal(getAllService.Body, &users); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
