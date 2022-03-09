package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DeleteUserHandle deletes User
func DeleteUserHandle() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// id, err := strconv.Atoi(ps.ByName("id"))
		// if err != nil {
		// 	http.Error(w, "Bad request", http.StatusBadRequest)
		// }

		// delete user ...
	}
}
