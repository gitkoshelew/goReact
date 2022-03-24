package middleware

import (
	"auth/domain/model"
	"auth/domain/store"
	"auth/internal/apperror"
	"auth/pkg/response"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ValidateLogin ...
func ValidateLogin(next http.Handler, s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		login := &model.Login{}
		if err := json.NewDecoder(r.Body).Decode(login); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Eror during JSON request decoding", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while opening auth DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err = login.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating login. Err msg: %v", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Error occured while validating login", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), store.LoginValidateCtXKey, login)))
	})
}
