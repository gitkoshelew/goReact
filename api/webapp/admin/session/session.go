package session

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

/*func SessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["login"] = "login"
	session.Values["password"] = "password"
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}*/

func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["accountID"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}
}

func AuthSession(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := store.Get(r, "session")
	session.Values["accountID"] = id
	session.Save(r, w)
}
