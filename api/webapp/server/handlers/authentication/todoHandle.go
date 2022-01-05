package authentication

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TodoHandle ...
func TodoHandle() httprouter.Handle {

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
