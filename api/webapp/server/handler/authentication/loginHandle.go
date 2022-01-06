package authentication

import (
	"encoding/json"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/server/handler/request"
	"log"
	"net/http"
)

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		s.Open()
		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		tk, err := CreateToken(uint64(user.UserID), string(user.Role))

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)
		w.Header().Add("Access-Token", tk.AccessToken)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)

	})
}
