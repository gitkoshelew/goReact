package server

import (
	handlersadmin "goReact/webapp/admin/handlersAdmin"
	accountHandlers "goReact/webapp/admin/handlersAdmin/accountHandlers"
	bookingHandlers "goReact/webapp/admin/handlersAdmin/bookingHandlers"
	employeeHandlers "goReact/webapp/admin/handlersAdmin/employeeHandlers"
	hotelhandlers "goReact/webapp/admin/handlersAdmin/hotelHandlers"
	pethandlers "goReact/webapp/admin/handlersAdmin/petHandlers"
	roomhandlers "goReact/webapp/admin/handlersAdmin/roomHandlers"
	seathandlers "goReact/webapp/admin/handlersAdmin/seatHandlers"	
	usershandlers "goReact/webapp/admin/handlersAdmin/usersHandlers"

	"net/http"
)

func (s *Server) configureRoutesAdmin() {
	s.router.Handle("GET", "/admin/login", handlersadmin.LoginAdmin())

	s.router.Handle("GET", "/admin/home", handlersadmin.HomeAdmin())

	s.router.Handle("GET", "/admin/users", usershandlers.AllUsersHandler())
	s.router.Handle("GET", "/admin/users/id:id", usershandlers.GetUserByID())

	s.router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler())
	s.router.Handle("GET", "/admin/hotels/id:id", hotelhandlers.GetHotelByID())

	s.router.Handle("GET", "/admin/pets", pethandlers.AllPetsHandler())
	s.router.Handle("GET", "/admin/pets/id:id", pethandlers.GetPetByID())

	s.router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler())
	s.router.Handle("GET", "/admin/rooms/id:id", roomhandlers.GetRoomByID())

	s.router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler())
	s.router.Handle("GET", "/admin/seats/id:id", seathandlers.GetSeatByID())

	s.router.Handle("GET", "/admin/accounts", accountHandlers.AllAccountsHandler())
	s.router.Handle("GET", "/admin/accounts/id:id", accountHandlers.GetAccountByID())

	s.router.Handle("GET", "/admin/employees", employeeHandlers.AllEmployeeHandler())
	s.router.Handle("GET", "/admin/employees/id:id", employeeHandlers.GetEmployeeByID())

	s.router.Handle("GET", "/admin/bookings", bookingHandlers.AllBookingsHandler())
	s.router.Handle("GET", "/admin/bookings/id:id", bookingHandlers.GetBookingByID())

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))

}
