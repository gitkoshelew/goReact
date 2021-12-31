package authentication

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// RefreshHandle ...
func RefreshHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		refreshToken := ExtractRefreshToken(r)

		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Refresh token expired", http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			role := fmt.Sprint(claims["role"])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			tk, createErr := CreateToken(userID, role)
			if createErr != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			c := http.Cookie{
				Name:     "JWT",
				Value:    tk.RefreshToken,
				HttpOnly: true,
			}

			http.SetCookie(w, &c)

			w.Header().Add("Access-Token", tk.AccessToken)
			w.WriteHeader(http.StatusOK)

		} else {
			http.Error(w, "refresh expired", http.StatusUnauthorized)
		}

	}
}
