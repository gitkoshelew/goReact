package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// EmailConfirm ...
func EmailConfirm(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		endp := ps.ByName("token")

		// token, err := jwt.Parse(endp,
		// 	func(token *jwt.Token) (interface{}, error) {
		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 			s.Logger.Errorf("Unexpected signing method. %v", token.Header["alg"])
		// 			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Unexpected signing method. %v", token.Header["alg"])})
		// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 		}
		// 		return []byte(EmailSecretKey), nil
		// 	})
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	s.Logger.Errorf("Link is expired. Errors msg: %v", err)
		// 	json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		// 	return
		// }

		// if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	s.Logger.Errorf("Can't parse link. Errors msg: %v", err)
		// 	json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		// 	return
		// }

		token, err := ParseCustomToken(endp, EmailSecretKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can't parse link. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		}

		id, err := strconv.Atoi(fmt.Sprint(token["user_id"]))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			s.Logger.Errorf("Eror while parsing token. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.User().VerifyEmail(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant update user. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Email confirms successfully!"})
	}

}
