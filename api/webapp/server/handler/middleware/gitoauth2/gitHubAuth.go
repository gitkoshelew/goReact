package gitoauth2

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



func GitHubAuth(next http.HandlerFunc, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		r.Header.Set("Accept", "application/json")

		code := r.URL.Query().Get("code")

		url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", conf.Endpoint.TokenURL, conf.ClientID, conf.ClientSecret, code)	

		token , err := service.GetTokenGit (url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while making resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}

		/*req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while making requsr . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}
		req.Header.Set("Accept", "application/json")
		var c http.Client
		resp, err := c.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while getting resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}
		defer resp.Body.Close()

		token := &oauth2.Token{}

		
		if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err ms:%v. Response body: %v", err, &resp.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}*/



		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CTXKeyAccessTokenGitOAuth, token)))

	}
}
