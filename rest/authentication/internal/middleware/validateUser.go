package middleware

import (
	"auth/domain/model"
	"auth/domain/store"
	"auth/internal/apperror"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ValidateUser ...
func ValidateUser(next http.Handler, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		userDTO := &model.UserDTO{}
		if err := json.NewDecoder(r.Body).Decode(userDTO); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(apperror.NewAppError("Eror during JSON request decoding", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while opening auth DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err = userDTO.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating user. Err msg: %v", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while validating user", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), store.UserValidateCtXKey, userDTO)))
	}

}
