package restorepassword

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// ChangePassword ...
func ChangePassword(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}

		email, err := handler.ContextEmail(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Cannot parse email: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Cannot parse email: %v", err)})
			return
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Cannot parse email: %v", err)})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Cannot open DB: Err msg: %v", err)})
			return
		}

		user, err := s.User().FindByEmail(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching user: %v", err)})
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
		if err == nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while searching user")
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching user: %v", err)})
			return
		}

		user.Password = req.Password

		u := s.User().ModelFromDTO(user)

		err = s.User().PasswordChange(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while searching user: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Password changed for user= %d", u.UserID)})
	}
}
