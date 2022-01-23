package usershandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteUser ...
func DeleteUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		err = s.User().Delete(id)
		if err != nil {
			log.Print(err)
			s.Logger.Errorf("Can't delete user. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Delete user with id = %d", id)
		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
