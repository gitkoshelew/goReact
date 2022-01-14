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
		s.Open()
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		role := r.FormValue("Role")
		verified, _ := strconv.ParseBool(r.FormValue("Verified"))
		name := r.FormValue("Name")
		surname := r.FormValue("Surname")
		middleName := r.FormValue("MiddleName")
		sex := r.FormValue("Sex")

		layout := "2006-01-02"
		dateOfBirth, _ := time.Parse(layout, r.FormValue("DateOfBirth"))
		address := r.FormValue("Address")
		phone := r.FormValue("Phone")
		photo := r.FormValue("Photo")

		u := model.User{
			UserID:      0,
			Email:       email,
			Password:    password,
			Role:        model.Role(role),
			Name:        name,
			Surname:     surname,
			MiddleName:  middleName,
			Sex:         model.Sex(sex),
			Address:     address,
			Phone:       phone,
			Photo:       photo,
			Verified:    verified,
			DateOfBirth: dateOfBirth,
		}

		err := u.NewUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = s.User().Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
