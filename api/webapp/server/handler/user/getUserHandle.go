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
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		err = s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user, err := s.User().FindByID(id)
		if err != nil {
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)
	}
}
