package handlers

import (
	"authentication/domain/model"
	"authentication/internal/apperror"
	"authentication/internal/store"
	jwthelper "authentication/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Login ...
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Eror during JSON request decoding. Request body: %v", r.Body), fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), err.Error()))
			return
		}
		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Invalid email or password", fmt.Sprintf("%d", http.StatusBadRequest), err.Error()))
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
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
	}
}
