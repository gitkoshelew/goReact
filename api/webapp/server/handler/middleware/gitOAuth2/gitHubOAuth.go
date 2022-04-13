package gitoauth2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

var (
	client_id     = os.Getenv("GIT_CLIENT_ID")
	client_secret = os.Getenv("GIT_CLIENT_SECRET")
	authorize_uri = os.Getenv("AUTHOTIZE_URI")
	redirect_uri  = os.Getenv("REDIRECT_URI")
)

func gitHubOAuth(w http.ResponseWriter, r *http.Request) {

	link := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", authorize_uri, client_id, redirect_uri)

	http.Redirect(w, r, link, http.StatusFound)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: client_secret},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	//repos, _, err := client.Repositories.List(ctx, "", nil)
	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))

}
