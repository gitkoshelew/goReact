package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		_, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("Unauthorized. Msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("You are unauthorized. Err msg: %v", err)})
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
