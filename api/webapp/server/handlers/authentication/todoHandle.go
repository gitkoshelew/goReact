package authentication

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Todo ...
type Todo struct {
	UserID uint64 `json:"userId"`
	Title  string `json:"title"`
}

// TodoHandle ...
func TodoHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		tokenAuth, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Print(err.Error())
			return
		}

		userID, err := FetchAuth(tokenAuth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Print(err.Error())
			return
		}

		var td Todo
		td.UserID = userID

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(td)
	}
}
