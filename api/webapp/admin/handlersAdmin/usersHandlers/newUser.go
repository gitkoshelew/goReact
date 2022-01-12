package usershandlers

import (
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"goReact/webapp/validation"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		s.Open()
		email := r.FormValue("Email")
		validation.ValidateEmail(email, w)

		password := r.FormValue("Password")

		role := r.FormValue("Role")

		verified, _ := strconv.ParseBool(r.FormValue("Verified"))

		name := r.FormValue("Name")

		err :=validation.ValidateName(name, w)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		surname := r.FormValue("Surname")
		err = validation.ValidateName(surname, w)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		middleName := r.FormValue("MiddleName")
		err = validation.ValidateName(middleName, w)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sex := r.FormValue("Sex")

		
		layout := "2006-01-02"
		dateOfBirth, _ := time.Parse(layout, r.FormValue("DateOfBirth"))
		validation.ValidateDateOfBirth(dateOfBirth, w)
		
		address := r.FormValue("Address")
		phone := r.FormValue("Phone")
		photo := r.FormValue("Photo")

		u := model.NewUser(
			0,
			email,
			password,
			role,
			name,
			surname,
			middleName,
			sex,
			address,
			phone,
			photo,
			verified,
			dateOfBirth,
		)

		_, err = s.User().Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
