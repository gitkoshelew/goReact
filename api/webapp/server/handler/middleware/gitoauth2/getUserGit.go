package gitoauth2

import (
	"encoding/json"
	"fmt"
	reqandresp "goReact/domain/reqAndResp"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var getUserURI = os.Getenv("GIT_GET_USER_DATA")

func GetUserGit(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Context().Value(handler.CTXKeyAccessTokenGitOAuth).(*oauth2.Token)

		authHeader := fmt.Sprintf("token %s", token.AccessToken)

		w.Header().Set("Accept", "application/json")

		result, err := service.GetUserData(getUserURI, authHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while getting resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}

		gitSSOUser := &reqandresp.GitSSOUser{}
		err = json.Unmarshal(*result, gitSSOUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred  while unmarshaling bytes. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		user, err := reqandresp.UserFromGit(gitSSOUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred  while converting user. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB. Err msg: %v", err)})
		}
		_, err = s.User().CreateFromSocial(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while creating user: %v", err)})
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
