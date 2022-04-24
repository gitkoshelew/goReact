package gitoauth2

import (
	"goReact/domain/store"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:     os.Getenv("GIT_CLIENT_ID"),
	ClientSecret: os.Getenv("GIT_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GIT_REDIRECT_URI"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  os.Getenv("GIT_AUTHOTIZE_URI"),
		TokenURL: os.Getenv("GIT_ACCESS_TOKEN_URI"),
	},
	
}

func GitHubLogin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		link := conf.AuthCodeURL("state" , )

		http.Redirect(w, r, link, http.StatusFound)

	}
}
