package handlersadmin

import (
	"goReact/webapp/admin/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutAdmin ...
func LogoutAdmin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.Logout(w, r)
		w.WriteHeader(http.StatusOK)
	}
}
