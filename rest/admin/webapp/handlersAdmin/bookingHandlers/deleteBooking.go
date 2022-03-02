package bookinghandlers

import (
	"admin/domain/model"
	"admin/domain/store"
	"admin/webapp/session"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_delete model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "delete_bookings",
	Descriptoin:  "ability to read a booking"}

// Deletebooking ...
func DeleteBooking(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_delete.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			http.Redirect(w, r, "/admin/homebookings", http.StatusFound)
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		err = s.Booking().Delete(id)
		if err != nil {
			log.Print(err)
			s.Logger.Errorf("Can't delete booking. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Delete booking with id = %d", id)
		http.Redirect(w, r, "/admin/homebookings", http.StatusFound)

	}
}
