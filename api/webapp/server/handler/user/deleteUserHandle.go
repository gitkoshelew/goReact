package user

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteUserHandle deletes User
func DeleteUserHandle(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while parsing id: %v", err)})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while opening DB: %v", err)})
			return
		}

		err = s.User().Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while deleting user: %v", err)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Userd with id %d was deleted", id)})
	}
}
