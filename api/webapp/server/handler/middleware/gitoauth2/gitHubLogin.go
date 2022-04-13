package gitoauth2

import (
	"fmt"
	"goReact/domain/store"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

var (
	client_id     = os.Getenv("GIT_CLIENT_ID")
	//authorize_uri = os.Getenv("GIT_AUTHORIZE_UI")
	redirect_uri  = os.Getenv("GIT_REDIRECT_URI")

	Endpoint = oauth2.Endpoint{
		AuthURL:  os.Getenv("AUTHOTIZE_URI"),
		TokenURL: os.Getenv("GIT_ACCESS_TOKEN_URI"),
	}
)

func GitHubLogin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", Endpoint.AuthURL, client_id, redirect_uri)

		http.Redirect(w, r, url, http.StatusFound)

	}
}
