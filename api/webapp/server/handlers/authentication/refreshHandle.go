package authentication

import (
	"encoding/json"
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

		type refreshRequets struct {
			RefreshToken string `json:"refreshToken"`
		}
		req := &refreshRequets{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		refreshToken := req.RefreshToken

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

			refreshUUID, ok := claims["refresh_uuid"].(string)
			if !ok {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			deleted, err := DeleteAuth(refreshUUID)

			if err != nil || deleted == 0 {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			ts, createErr := CreateToken(userID)
			if createErr != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			err = CreateAuth(userID, ts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			tokens := map[string]string{
				"accessToken":  ts.AccessToken,
				"refreshToken": ts.RefreshToken,
			}
			json.NewEncoder(w).Encode(tokens)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "refresh expired", http.StatusUnauthorized)
		}

	}
}
