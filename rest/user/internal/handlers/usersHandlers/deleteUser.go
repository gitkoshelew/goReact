package usershandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// DeleteUser ...
func DeleteUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Delete user with id = %d", id)})

	}
}
