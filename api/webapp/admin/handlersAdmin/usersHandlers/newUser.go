package usershandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var permissionCreat model.Permission = model.Permission{Name: model.CreatUser}

// NewUser ...
func NewUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionCreat.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.FormValue("Email")
		password := r.FormValue("Password")

		role := r.FormValue("Role")
		verified, err := strconv.ParseBool(r.FormValue("Verified"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified")), http.StatusBadRequest)
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
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
			return
		}
		address := r.FormValue("Address")
		phone := r.FormValue("Phone")
		photo := r.FormValue("Photo")

		userDTO := model.UserDTO{
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

		/*err = userDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}*/

		user := s.User().ModelFromDTO(&userDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v.", err)
			return
		}

		_, err = s.User().Create(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while creating user. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/homeusers/", http.StatusFound)

	}
}
