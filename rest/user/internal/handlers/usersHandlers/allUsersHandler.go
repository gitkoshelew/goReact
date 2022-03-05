package usershandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/internal/apperror"
	"user/internal/store"

	"github.com/julienschmidt/httprouter"
)

// AllUsersHandler ...
func AllUsersHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", r.Body), fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		users, err := s.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find users", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find users. Err msg: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
