package auth

import (
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type user struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
}

// LoginHandle ...
func LoginHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		loginService, err := auth.Login(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		var user *user
		if err := json.Unmarshal(loginService.Body, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		cookies := loginService.Cookies
		for _, cookie := range cookies {
			http.SetCookie(w, cookie)
		}
		w.Header().Set("Access-Token", loginService.Headers["Access-Token"])
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
