package user

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetUsersHandle returns all Users
func GetUsersHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM USERS")
		if err != nil {
			log.Fatal(err)
		}

		usersDto := []dto.UserDto{}

		for rows.Next() {
			user := dto.UserDto{}
			err := rows.Scan(
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
				log.Printf(err.Error())
				continue
			}

			usersDto = append(usersDto, user)
		}

		json.NewEncoder(w).Encode(usersDto)
	}
}
