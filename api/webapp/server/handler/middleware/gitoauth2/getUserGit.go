package gitoauth2

import (
	"context"
	"encoding/json"
	"fmt"

	"goReact/domain/reqAndResp/oauth"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler"

	"goReact/webapp/server/handler/response"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var getUserURI = os.Getenv("GIT_GET_USER_DATA")

//Getting a resource (user) and creating a user
func GetUserGit(next http.HandlerFunc, s *store.Store) http.HandlerFunc {
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

		gitUser := &oauth.GitSSOUser{}
		err = json.Unmarshal(*result, gitUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred  while unmarshaling bytes. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		fmt.Println(gitUser.Email)
		fmt.Println(gitUser.UserID)
		fmt.Println(gitUser.Name)


		err = gitUser.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while validating user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while validating user. Err msg: %v", err)})
			return
		}

		user, err := oauth.UserFromGit(gitUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred  while converting user. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		_, err = service.OauthRegisterUser(s, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred while registering user. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handler.CtxKeyUser, user)))

	}
}
