package linkedinoauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/service"

	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LinkedInAuth(next http.HandlerFunc, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		r.Header.Set("Accept", "application/json")

		code := r.URL.Query().Get("code")

		url := fmt.Sprintf("%s?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s",
			conf.Endpoint.TokenURL,
			code,
			conf.RedirectURL,
			conf.ClientID,
			conf.ClientSecret,
		)

		token, err := service.GetToken(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while making resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CTXKeyAccessTokenLinkedINOAuth, token)))

	}
}
