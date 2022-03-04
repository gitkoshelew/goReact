package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"gateway/internal/client/customer"
	"gateway/internal/client/hotel"
	"gateway/pkg/response"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		rBody, _ := ioutil.ReadAll(r.Body)

		_, err := auth.Registration(r.Context(), service, bytes.NewReader(rBody))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		user := &UserDTO{}
		if err := json.NewDecoder(bytes.NewReader(rBody)).Decode(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			service.Base.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}
		user.Verified = false

		switch user.Role {
		case "client":
			_, err := customer.Create(r.Context(), client.CustomerUserService, bytes.NewReader(rBody))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
				return
			}
		case "employee":
			_, err := hotel.Create(r.Context(), client.HotelEmployeeService, bytes.NewReader(rBody))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "User created!"})

	}
}
