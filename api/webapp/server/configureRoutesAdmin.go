package server

import (
	handlersadmin "goReact/webapp/admin/handlersAdmin"
	bookingHandlers "goReact/webapp/admin/handlersAdmin/bookingHandlers"
	employeeHandlers "goReact/webapp/admin/handlersAdmin/employeeHandlers"
	hotelhandlers "goReact/webapp/admin/handlersAdmin/hotelHandlers"
	pethandlers "goReact/webapp/admin/handlersAdmin/petHandlers"
	roomhandlers "goReact/webapp/admin/handlersAdmin/roomHandlers"
	usershandlers "goReact/webapp/admin/handlersAdmin/usersHandlers"

	"net/http"
)

func (s *Server) configureRoutesAdmin() {
	s.router.HandlerFunc("GET", "/admin/login", handlersadmin.LoginAdmin())

	s.router.HandlerFunc("GET", "/admin/home", handlersadmin.HomeAdmin())

	s.router.HandlerFunc("GET", "/admin/allusers", usershandlers.AllUsersHandler())

	s.router.HandlerFunc("GET", "/admin/allhotels", hotelhandlers.AllHotelsHandler())

	s.router.HandlerFunc("GET", "/admin/allpets", pethandlers.AllPetsHandler())

	s.router.HandlerFunc("GET", "/admin/allrooms", roomhandlers.AllRoomsHandler())

	s.router.HandlerFunc("GET", "/admin/allemployees", employeeHandlers.AllEmployeeHandler())

	s.router.HandlerFunc("GET", "/admin/allbookings", bookingHandlers.AllBookingsHandler())

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))

}
