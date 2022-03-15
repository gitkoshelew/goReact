package customer

import (
	"context"
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/customer"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetUserHandle ...
func GetUserHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		getService, err := customer.Get(context.WithValue(r.Context(), client.CustomerGetQuerryParamsCtxKey, ps.ByName("id")), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var user *UserDTO
		if err := json.Unmarshal(getService.Body, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
