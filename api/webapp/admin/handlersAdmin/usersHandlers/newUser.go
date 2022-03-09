package usershandlers

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// NewUser ...
func NewUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		role := r.FormValue("Role")
		verified, err := strconv.ParseBool(r.FormValue("Verified"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified"))
			return
		}
		name := r.FormValue("Name")
		surname := r.FormValue("Surname")
		middleName := r.FormValue("MiddleName")
		sex := r.FormValue("Sex")

		layout := "2006-01-02"
		dateOfBirth, err := time.Parse(layout, r.FormValue("DateOfBirth"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
			return
		}
		address := r.FormValue("Address")
		phone := r.FormValue("Phone")
		photo := r.FormValue("Photo")

		u := model.UserDTO{
			UserID:      0,
			Email:       email,
			Password:    password,
			Role:        role,
			Name:        name,
			Surname:     surname,
			MiddleName:  middleName,
			Sex:         sex,
			Address:     address,
			Phone:       phone,
			Photo:       photo,
			Verified:    verified,
			DateOfBirth: dateOfBirth,
		}
		_, err = s.User().Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create user. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat user with id = %d", u.UserID)
		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
