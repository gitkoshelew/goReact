package user

import (
	"encoding/json"
	"goReact/domain/store"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetUserHandle returns User by ID
func GetUserHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		row := db.QueryRow("SELECT * FROM USERS WHERE id = $1", id)

		user := store.User{}
		err = row.Scan(
			&user.UserID,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Account)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(dto.UserToDto(user))
	}
}
