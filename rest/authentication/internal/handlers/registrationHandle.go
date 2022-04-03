package handlers

import (
	"auth/domain/model"
	"auth/domain/store"
	"auth/internal/apperror"
	"auth/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegistrationHandle ...
func RegistrationHandle(s *store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userDTO := r.Context().Value(store.UserValidateCtXKey).(*model.UserDTO)

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		user, err := s.User().ModelFromDTO(userDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while building user model from DTO. Err msg:%v.", err)
			return
		}

		_, err = s.User().Create(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Cant create user", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", user.UserID)})
	})
}
