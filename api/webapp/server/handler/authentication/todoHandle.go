package authentication

import (
	"encoding/json"
	"goReact/domain/store"
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
			s.Logger.Errorf("Can't. Errors msg: %v", err)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		response := accessDetailsResponse{
			UserID: AccessDetails.UserID,
			Role:   AccessDetails.Role,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
