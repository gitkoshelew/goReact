package authentication

import (
	"encoding/json"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		login := r.Context().Value(handler.CtxKeyLoginValidation).(*model.Login)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		user, err := s.User().FindByEmail(login.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.CheckPasswordHash(user.Password, login.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		tk, err := CreateToken(uint64(user.UserID), string(user.Role))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)
		w.Header().Set("Access-Token", tk.AccessToken)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})
}
