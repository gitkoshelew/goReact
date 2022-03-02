package usershandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permission_creat model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "creat_user",
	Descriptoin:  "ability to creat a user"}

// NewUser ...
func NewUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_creat.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
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

		err = u.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = u.WithEncryptedPassword()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v.", err)
			return
		}

		_, err = s.User().Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't create user. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Creat user with id = %d", u.UserID)
		http.Redirect(w, r, "/admin/homeusers/", http.StatusFound)

	}
}
