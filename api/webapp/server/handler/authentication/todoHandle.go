package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TodoHandle ...
func TodoHandle(s *store.Store) httprouter.Handle {

	type accessDetailsResponse struct {
		UserID uint64 `json:"userId"`
		Role   string `json:"role"`
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		AccessDetails, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Print(err.Error())
			fmt.Fprintln(w, "You're not authorized")
			return
		}

		response := accessDetailsResponse{
			UserID: AccessDetails.UserID,
			Role:   AccessDetails.Role,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
