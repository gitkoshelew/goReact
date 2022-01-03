package middleware

import (
	"context"
	"goReact/webapp/server/handlers/authentication"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthenticateUser ...
func AuthenticateUser(next http.Handler) httprouter.Handle {
	db := utils.HandlerDbConnection()
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		AccessDetails, err := authentication.ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Print(err.Error())
			return
		}

		row := db.QueryRow("SELECT * FROM USERS WHERE id = $1", AccessDetails.UserID)

		user := dto.UserDto{}
		err = row.Scan(
			&user.UserID,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&user.DateOfBirth,
			&user.Address,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Verified,
			&user.Sex,
			&user.Photo,
		)
		if err != nil {
			panic(err)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyUser, user)))
	})

}
