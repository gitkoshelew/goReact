package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/authentication"
	"net/http"
)

// AuthenticateUser ...
func AuthenticateUser(next http.Handler, s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AccessDetails, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			s.Logger.Errorf("Extract token meta data error. MSG: %v", err)
			json.NewEncoder(w).Encode(fmt.Sprintf("You are unauthorized. Err msg: %v", err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		s.Open()
		user, err := s.User().FindByID(int(AccessDetails.UserID))
		if err != nil {
			s.Logger.Errorf("Can't find user. Error msg: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyUser, user)))
	})

}
