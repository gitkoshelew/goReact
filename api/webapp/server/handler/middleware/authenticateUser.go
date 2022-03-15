package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/authentication"
	"goReact/webapp/server/handler/response"
	"net/http"
)

// AuthenticateUser ...
func AuthenticateUser(next http.Handler, s *store.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AccessDetails, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while token extraction. Err msg: %v", err)})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while DB opennig. Err msg: %v", err)})
			return
		}

		user, err := s.User().FindByID(int(AccessDetails.UserID))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while user searching. Err msg: %v", err)})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CtxKeyUser, user)))
	})

}
