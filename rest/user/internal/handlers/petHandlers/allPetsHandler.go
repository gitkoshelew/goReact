package pethandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/internal/apperror"
	"user/internal/store"

	"github.com/julienschmidt/httprouter"
)

// AllPetsHandler ...
func AllPetsHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}
		pets, err := s.Pet().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find pets", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find pets. Err msg: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pets)

	}
}
