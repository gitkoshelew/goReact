package permission

import (
	"fmt"
	"goReact/domain/model"
	"goReact/domain/store"
	"goReact/webapp/admin/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// AddPermissionsEmployee ...
func AddPermissionsEmployee(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, "admin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Access is denied. Err msg:%v. ", err)
			return
		}

		employeeID, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		perID, err := strconv.Atoi(r.FormValue("permissions"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("permissions")), http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("permissions"))
			return
		}
		peDTO := model.PermissionsEmployeesDTO{
			PermissionsID: perID,
			EmployeeID:    employeeID,
		}

		err = s.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = peDTO.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		pe, err := s.PermissionsEmployee().ModelFromDTO(&peDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.PermissionsEmployee().SetForEmployee(pe)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while setting permissions. Err msg:%v. ", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/addpermissions", http.StatusFound)

	}
}
