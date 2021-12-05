package user

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteUserHandle deletes User
func DeleteUserHandle() httprouter.Handle {

	users := dto.GetUsersDto()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		for index, u := range users {
			if u.UserID == id { // delete object imitation =)
				users[index].Name = "DELETE"
				users[index].Surname = "DELETE"
				users[index].MiddleName = "DELETE"
				json.NewEncoder(w).Encode(users)
				return
			}
		}

		http.Error(w, "Cant find User", http.StatusBadRequest)
	}
}
