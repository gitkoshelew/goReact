package user

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetUsersHandle returns all users
func GetUsersHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB: %v", err)})
			return
		}

		users, err := s.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't find user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting all users DB: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
