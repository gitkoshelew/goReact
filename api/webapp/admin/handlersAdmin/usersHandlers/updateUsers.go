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

var permissionUpdate model.Permission = model.Permission{Name: model.UpdateUser}

// UpdateUser ...
func UpdateUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permissionUpdate.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userDTO, err := s.User().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := r.FormValue("Email")
		if email != "" {
			userDTO.Email = email
		}

		verified := r.FormValue("Verified")
		if verified != "" {
			verified, err := strconv.ParseBool(verified)
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("Verified"))
				return
			}
			userDTO.Verified = verified
		}

		role := r.FormValue("Role")
		if role != "" {
			userDTO.Role = role
		}

		name := r.FormValue("Name")
		if name != "" {
			userDTO.Name = name
		}

		surname := r.FormValue("Surname")
		if surname != "" {
			userDTO.Surname = surname
		}

		middleName := r.FormValue("MiddleName")
		if surname != "" {
			userDTO.MiddleName = middleName
		}

		sex := r.FormValue("Sex")
		if sex != "" {
			userDTO.Sex = sex
		}

		layout := "2006-01-02"
		dateOfBirth := r.FormValue("DateOfBirth")
		if dateOfBirth != "" {
			dateOfBirth, err := time.Parse(layout, r.FormValue("DateOfBirth"))
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth")), http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
				return
			}
			userDTO.DateOfBirth = dateOfBirth
		}

		address := r.FormValue("Address")
		if address != "" {
			userDTO.Address = address
		}

		phone := r.FormValue("Phone")
		if phone != "" {
			userDTO.Phone = phone
		}

		photo := r.FormValue("Photo")
		if photo != "" {
			userDTO.Photo = photo
		}

		/*err = u.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}*/

		user := s.User().ModelFromDTO(userDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v.", err)
			return
		}

		err = s.User().Update(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while updating user. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/homeusers", http.StatusFound)

	}
}
