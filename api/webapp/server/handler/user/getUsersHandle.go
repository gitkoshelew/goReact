package user

import (
	"encoding/json"
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
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)

		}

		users, err := s.User().GetAll()
		if err != nil {
			s.Logger.Errorf("Can't find user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
