package authentication

import (
	"encoding/json"
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler/request"
	"goReact/webapp/server/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(s *store.Store, m *service.Mail) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &request.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		u := model.User{
			UserID:      0,
			Email:       req.Email,
			Password:    req.Password,
			Role:        model.Role(req.Role),
			Name:        req.Name,
			Surname:     req.Surname,
			MiddleName:  req.MiddleName,
			Sex:         model.Sex(req.Sex),
			Address:     req.Address,
			Phone:       req.Phone,
			Photo:       req.Photo,
			Verified:    false,
			DateOfBirth: req.DateOfBirth,
		}

		err := u.WithEncryptedPassword()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		_, err = s.User().Create(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Cant create user. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		payload := make(map[string]string)
		payload["user_id"] = strconv.Itoa(u.UserID)

		endpoint, err := CreateCustomToken(payload, 120, EmailSecretKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Error while email confirmation endpoint creating. Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		m.Send(service.EmailConfirmation, endpoint, []string{u.Email})
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("User id = %d", u.UserID)})
	}
}
