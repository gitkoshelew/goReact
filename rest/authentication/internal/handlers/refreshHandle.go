package handlers

import (
	"auth/domain/store"
	"auth/internal/apperror"
	jwthelper "auth/pkg/jwt"
	"auth/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/julienschmidt/httprouter"
)

// RefreshHandle ...
func RefreshHandle(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		refreshToken := jwthelper.ExtractRefreshToken(r)

		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				s.Logger.Errorf("Unexpected signing method. %v", token.Header["alg"])
				json.NewEncoder(w).Encode(json.NewEncoder(w).Encode(apperror.NewAppError("Unexpected signing method", fmt.Sprintf("%d", http.StatusUnprocessableEntity), fmt.Sprintf("%v", token.Header["alg"]))))
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("refresh token expired. err: %w", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("refresh token expired", fmt.Sprintf("%d", http.StatusUnauthorized), err.Error()))
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			w.WriteHeader(http.StatusUnprocessableEntity)
			s.Logger.Errorf("can't parse token. err: %w", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("can't parse token", fmt.Sprintf("%d", http.StatusUnprocessableEntity), err.Error()))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("eror while parsing token. err: %w", err)
				json.NewEncoder(w).Encode(apperror.NewAppError("eror while parsing token", fmt.Sprintf("%d", http.StatusUnprocessableEntity), err.Error()))
				return
			}

			role := fmt.Sprint(claims["role"])
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("eror while parsing token. err: %w", err)
				json.NewEncoder(w).Encode(apperror.NewAppError("eror while parsing token", fmt.Sprintf("%d", http.StatusUnprocessableEntity), err.Error()))
				return
			}

			tk, createErr := jwthelper.CreateToken(userID, role)
			if createErr != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(apperror.NewAppError("eror while parsing token", fmt.Sprintf("%d", http.StatusUnprocessableEntity), err.Error()))
				return
			}

			c := http.Cookie{
				Name:     "Refresh-Token",
				Value:    tk.RefreshToken,
				HttpOnly: true,
			}

			http.SetCookie(w, &c)
			w.Header().Add("Access-Token", tk.AccessToken)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully refreshed"})

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Error("refresh token is expired")
			json.NewEncoder(w).Encode(apperror.NewAppError("refresh token is expired", fmt.Sprintf("%d", http.StatusUnauthorized), err.Error()))
		}

	}
}
