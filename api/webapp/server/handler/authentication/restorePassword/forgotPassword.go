package restorePassword

import (
	"encoding/json"
	"fmt"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// ForgotPassword pass ...
func ForgotPassword(s *store.Store, m *service.Mail) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		u, err := s.User().FindByEmail(req.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf(" No user with this email %s", err.Error())
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		epClaims := jwt.MapClaims{}
		epClaims["user_id"] = u.UserID
		epClaims["user_email"] = u.Email
		epClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
		
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, epClaims)
		endpoint, err := at.SignedString([]byte(os.Getenv("RESTORE_PASSWORD_SECRET")))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Error while password reset endpoint creating. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		m.Send(service.PassReset, endpoint, []string{u.Email})
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Messege sent = %s", u.Email)})
	}
}
