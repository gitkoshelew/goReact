package linkedinoauth2

import (
	"goReact/domain/store"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
	ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("LINKEDIN_REDIRECT_URI"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  os.Getenv("LINKEDIN_AUTHOTIZE_URI"),
		TokenURL: os.Getenv("LINKEDIN_ACCESS_TOKEN_URI"),
	},
	Scopes: []string{"r_emailaddress", "r_liteprofile"},
}

//Resource authorization request
func LinkedInLogin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		link := conf.AuthCodeURL("state")

		http.Redirect(w, r, link, http.StatusFound)

	}
}
