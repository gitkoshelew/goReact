package user

import (
	"encoding/json"
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

		user := dto.UserDto{}
		err = row.Scan(
			&user.UserID,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Sex,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Photo,
		)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(user)
	}
}
