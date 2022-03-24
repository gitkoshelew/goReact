package handlers

import (
	"auth/domain/model"
	"auth/domain/store"
	"auth/domain/utils"
	"auth/internal/apperror"
	jwthelper "auth/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
)

// Login ...
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		login := r.Context().Value(store.LoginValidateCtXKey).(*model.Login)
		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}
		user, err := s.User().FindByEmail(login.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Invalid email or password", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		err = utils.CheckPasswordHash(user.Password, login.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Invalid email or password", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		tk, err := jwthelper.CreateToken(uint64(user.UserID), string(user.Role))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(apperror.NewAppError("Eror during createing tokens", fmt.Sprintf("%d", http.StatusUnprocessableEntity), err.Error()))
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)

		w.Header().Set("Access-Token", tk.AccessToken)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})
}
