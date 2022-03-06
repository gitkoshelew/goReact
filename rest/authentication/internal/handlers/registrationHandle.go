package handlers

import (
	"authentication/domain/model"
	"authentication/internal/apperror"
	"authentication/internal/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type authDataID struct {
	ID int `json:"authDataId,omitempty"`
}

// RegistrationHandle ...
func RegistrationHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(apperror.NewAppError("Bad request", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		err := req.WithEncryptedPassword()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("bad request. requests body: %v", r.Body), fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		user, err := s.User().Create(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Cant create user", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}
		authData := authDataID{
			ID: user.UserID,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(authData)
	}
}
