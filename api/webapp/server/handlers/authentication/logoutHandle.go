package authentication

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		_, err := ExtractTokenMetadata(r)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		c := http.Cookie{
			Name:     "JWT",
			Value:    "",
			HttpOnly: true,
		}
		w.Header().Add("Access-Token", "")
		http.SetCookie(w, &c)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Successfully logged out")
	}
}
