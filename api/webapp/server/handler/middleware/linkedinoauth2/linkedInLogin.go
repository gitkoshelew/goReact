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
	Scopes: []string{"r_liteprofile" , "r_emailaddress"},
}

func LinkedInLogin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		link := conf.AuthCodeURL("state")
		s.Logger.Info("link", link)

		http.Redirect(w, r, link, http.StatusFound)

	}
}

//https://www.linkedin.com/oauth/v2/authorization?client_id=781ygi34xfb27y&redirect_uri=http%!A(string=&scope=r_liteprofile%20r_emailaddress%20)%!F(MISSING)%!F(MISSING)localhost%!A(MISSING)8080%!F(MISSING)api%!F(MISSING)linkedinlogin%!F(MISSING)re&response_type=code&state=state
