package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		_, err := ExtractTokenMetadata(r)
		if err != nil {
			s.Logger.Errorf("Unauthorized. Msg: %v", err)
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			json.NewEncoder(w).Encode(fmt.Sprintf("You are unauthorized. Err msg: %v", err))
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    "",
			HttpOnly: true,
		}
		w.Header().Add("Access-Token", "")
		http.SetCookie(w, &c)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Successfully logged out")
	}
}
