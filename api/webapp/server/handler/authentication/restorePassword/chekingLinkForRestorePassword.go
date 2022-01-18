package restorePassword

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// chekingLingForRestorePassword ...
func СhekingLinkForRestorePassword(s *store.Store , next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		endp := ps.ByName("token")

		token, err := jwt.Parse(endp, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				s.Logger.Errorf("Unexpected signing method. %v", token.Header["alg"])
				json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Unexpected signing method. %v", token.Header["alg"])})
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("RESTORE_PASSWORD_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Link is expired. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't parse link. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		var email string
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email = fmt.Sprint(claims["user_email"])
			if email != "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("Eror while parsing token. Errors msg: %v", err)
				json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
				return
			}
		}

		/*_, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("Eror while parsing token. Errors msg: %v", err)
				json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
				return
			}
		}*/

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Link is confirmed"})

		//w.Header().Add("Token", endp)
		w.Header().Add("UserEmail", email)
		next.ServeHTTP(w, r)
	}
}
