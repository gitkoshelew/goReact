package gitoauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang.org/x/oauth2"
)

/*var (
	clientsecret    = os.Getenv("GIT_CLIENT_SECRET")
	//access_token_uri = os.Getenv("GIT_ACCESS_TOKEN_RI")
)*/

func GitHubAuth(next http.HandlerFunc, s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Accept", "application/json")

		code := r.URL.Query().Get("code")

		url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", conf.Endpoint.TokenURL, conf.ClientID, conf.ClientSecret, code)

		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while getting resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}
		defer resp.Body.Close()
		s.Logger.Info("resp.Body ....1 - ", resp.Body)

		token := oauth2.Token{}
		s.Logger.Info("token ....2- ", token)

		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Cannot parse sakura response: %v\n", err)
			return
		}

		/*if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err ms:%v. Response body: %v", err, resp.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}*/

		err = json.Unmarshal(resBody, &token)
		if err != nil {
			log.Printf("error decoding sakura response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", resBody)

		}
		/*json.NewDecoder(resp.Body).Decode(token)
		s.Logger.Info("token 1.1 -  ", token)

		json.NewEncoder(w).Encode(token)*/

		s.Logger.Info("resBody1 .... 3- ", resBody)
		s.Logger.Info("token .... ", token)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CTXKeyAccessTokenGitOAuth, token)))
		s.Logger.Info("resBody2 ....4- ")

	}
}
