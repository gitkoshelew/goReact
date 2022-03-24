package handlers

import (
	"auth/domain/store"
	"auth/internal/apperror"
	jwthelper "auth/pkg/jwt"
	"auth/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		s.Logger.Debugf("Auth token is: %v", r.Header.Get("Authorization"))

		_, err := jwthelper.ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(apperror.NewAppError("you are unauthorized", fmt.Sprintf("%d", http.StatusUnauthorized), err.Error()))
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    "",
			HttpOnly: true,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Access-Token", "")
		http.SetCookie(w, &c)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully logged out"})
	}
}
