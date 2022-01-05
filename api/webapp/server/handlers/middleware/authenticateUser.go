package middleware

import (
	"context"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/authentication"
	"log"
	"net/http"
)

// AuthenticateUser ...
func AuthenticateUser(next http.Handler, s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AccessDetails, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			fmt.Fprintln(w, "You are not authorized")
			w.WriteHeader(http.StatusUnauthorized)
			log.Print(err.Error())
			return
		}

		s.Open()
		user, err := s.User().FindByID(int(AccessDetails.UserID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyUser, user)))
	})

}
