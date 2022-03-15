package auth

import (
	"context"
	"encoding/json"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		logoutService, err := auth.Logout(context.WithValue(r.Context(), client.AccessTokenCtxKey, r.Header.Get("Authorization")), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		cookies := logoutService.Cookies
		for _, cookie := range cookies {
			http.SetCookie(w, cookie)
		}
		w.Header().Set("Access-Token", logoutService.Headers["Access-Token"])
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully logged out"})
	}
}
