package session

import (
	"admin/domain/model"
	"encoding/gob"
	"fmt"
	"net/http"
	"strings"
)

// CheckSession ...
func CheckSession(w http.ResponseWriter, r *http.Request) {

	session, err := sstore.PGStore.Get(r, "session-key")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["EmployeeID"]

	if !ok {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}
}

// AuthSession ...
func AuthSession(w http.ResponseWriter, r *http.Request, employee *model.Employee, permissions *[]model.Permission) {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	gob.Register(model.Employee{})
	session.Values["Employee"] = employee
	session.Values["EmployeeID"] = employee.EmployeeID
	session.Values["EmployeePosition"] = employee.PositionString()

	gob.Register([]model.Permission{})
	session.Values["Permissions"] = permissions

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

// IsExist ...
func IsExist(w http.ResponseWriter, r *http.Request) bool {

	session, err := sstore.PGStore.Get(r, "session-key")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	_, ok := session.Values["EmployeeID"]

	return ok
}

//CheckRigths of employee and return err if not enough rights
func CheckRigths(w http.ResponseWriter, r *http.Request, name string) error {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	position, ok := session.Values["EmployeePosition"]
	if !ok {
		err = fmt.Errorf("no permissions in session")
		return err
	}

	if position.(string) == "admin" {
		return nil
	}

	permissions, ok := session.Values["Permissions"]
	if !ok {
		err = fmt.Errorf("no permissions in session")
		return err
	}

	str := fmt.Sprintf("%v", permissions)

	contain := strings.Contains(str, name)
	if !contain {
		err = fmt.Errorf("not enough rights")
		return err
	}
	return nil
}

func IsAdmin(w http.ResponseWriter, r *http.Request) error {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	position, ok := session.Values["EmployeePosition"]
	if !ok {
		err = fmt.Errorf("no permissions in session")
		return err
	}

	if position.(string) != "admin" {
		err = fmt.Errorf("you are not admin")
		return err
	}

	return nil
}
