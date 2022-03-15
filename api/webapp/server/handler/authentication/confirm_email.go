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

		token, err := ParseCustomToken(endp, EmailSecretKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
		}

		id, err := strconv.Atoi(fmt.Sprint(token["user_id"]))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			s.Logger.Errorf("Eror occured while converting token userid from string to int. Errors msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.User().VerifyEmail(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Email confirms successfully!"})
	}

}
