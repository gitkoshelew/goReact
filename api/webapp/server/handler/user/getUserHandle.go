package user

import (
	"encoding/json"
	"goReact/domain/store"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetUserHandle returns User by ID
func GetUserHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		s.Open()
		user, err := s.User().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
