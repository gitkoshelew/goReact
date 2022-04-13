package gitoauth2

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"

	"golang.org/x/oauth2"
)

var getUserURI = "https://api.github.com/user"

func GetUserGit(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Context().Value(handler.CTXKeyAccessTokenGitOAuth).(*oauth2.Token)
		s.Logger.Info("token.AccessToken 11  -  ", token.AccessToken)

		autHeader := fmt.Sprintf("token %s", token.AccessToken)
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Authorization", autHeader)
		s.Logger.Info("token 12 -  ", autHeader)

		resp, err := http.Get(getUserURI)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while getting resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}
		defer resp.Body.Close()
		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)
		s.Logger.Info("token 13 -  ", result)

		json.NewEncoder(w).Encode(result)
	}
}
