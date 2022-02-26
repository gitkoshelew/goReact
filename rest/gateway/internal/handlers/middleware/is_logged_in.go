package middleware

import (
	"encoding/json"
	jwthelper "gateway/pkg/jwt"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IsLoggedIn ...
func IsLoggedIn(next httprouter.Handle) httprouter.Handle {

	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		_, err := jwthelper.ExtractTokenMetadata(r)
		if err != nil {
			next(w, r, ps)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.Info{Messsage: "You are already logged in."})
		return
	})
}
