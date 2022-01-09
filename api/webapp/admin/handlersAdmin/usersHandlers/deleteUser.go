package usershandlers

import (
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func DeleteUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		//id, _ := strconv.Atoi(ps.ByName("id"))
		id , _:= strconv.Atoi(r.FormValue("id"))
		s.Open()
		err := s.User().Delete(id)
		if err != nil {
			log.Print(err)
		}
		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
