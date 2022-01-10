package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["UserID"]
	if !ok {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}
}

func AuthSession(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := store.Get(r, "session")
	session.Values["UserID"] = id
	session.Save(r, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "UserID")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

func IsExist(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["UserID"]
	return ok
}
