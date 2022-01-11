package authentication

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// MeHandle ...
func MeHandle(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		AccessDetails, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("Can't. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.Open()
		if err != nil {
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		user, err := s.User().FindByID(int(AccessDetails.UserID))
		if err != nil {
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
