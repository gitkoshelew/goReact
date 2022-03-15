package auth

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		RegistrationService, err := auth.Registration(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var user *user
		if err := json.Unmarshal(RegistrationService.Body, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "User created!"})

	}
}
