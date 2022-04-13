package gitoauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"golang.org/x/oauth2"
)

var (
	client_secret    = os.Getenv("GIT_CLIENT_SECRET")
	//access_token_uri = os.Getenv("GIT_ACCESS_TOKEN_URI")
)

func GitHubAuth(next http.HandlerFunc, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Accept", "application/json")

		code := r.URL.Query().Get("code")

		url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", Endpoint.TokenURL, client_id, client_secret, code)

		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}
		defer resp.Body.Close()

		token := &oauth2.Token{}
		if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CTXKeyAccessTokenGitOAuth, token)))

	}
}
