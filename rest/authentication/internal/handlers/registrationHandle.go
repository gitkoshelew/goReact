package handlers

import (
	"authentication/domain/model"
	"authentication/internal/apperror"
	"authentication/internal/store"
	"authentication/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// User ...
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
}

// RegistrationHandle ...
func RegistrationHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(apperror.NewAppError("Bad request", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		u := model.User{
			UserID:   0,
			Email:    req.Email,
			Password: req.Password,
			Role:     model.Role(req.Role),
			Verified: false,
		}

		err := u.WithEncryptedPassword()
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

		_, err = s.User().Create(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Cant create user", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", u.UserID)})
	}
}
