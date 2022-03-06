package usershandlers

import (
	"customer/internal/apperror"
	"customer/internal/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AllUsersHandler ...
func AllUsersHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", r.Body),
				fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		users, err := s.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while getting all users",
				fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
