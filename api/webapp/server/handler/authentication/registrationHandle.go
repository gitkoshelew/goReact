package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(s *store.Store, m *service.Mail) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		user := &model.UserDTO{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		_, err = s.User().Create(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		payload := map[string]interface{}{
			"user_id": user.UserID,
		}

		endpoint, err := CreateCustomToken(payload, 120, EmailSecretKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		m.Send(service.EmailConfirmation, endpoint, []string{user.Email})
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", user.UserID)})
	}
}
