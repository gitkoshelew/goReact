package server

import (
	"goReact/domain/store"
	handlersadmin "goReact/webapp/admin/handlersAdmin"
	bookinghandlers "goReact/webapp/admin/handlersAdmin/bookingHandlers"
	employeehandlers "goReact/webapp/admin/handlersAdmin/employeeHandlers"
	hotelhandlers "goReact/webapp/admin/handlersAdmin/hotelHandlers"
	pethandlers "goReact/webapp/admin/handlersAdmin/petHandlers"
	roomhandlers "goReact/webapp/admin/handlersAdmin/roomHandlers"
	seathandlers "goReact/webapp/admin/handlersAdmin/seatHandlers"
	usershandlers "goReact/webapp/admin/handlersAdmin/usersHandlers"
	"net/http"
)

func (s *Server) configureRoutesAdmin() {
	s.router.Handle("GET", "/admin/login", handlersadmin.LoginAdmin(store.New(s.config)))
	s.router.Handle("POST", "/admin/auth", handlersadmin.AuthAdmin(store.New(s.config)))
	s.router.Handle("GET", "/admin/logout", handlersadmin.LogoutAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/home", handlersadmin.HomeAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/users", usershandlers.AllUsersHandler(store.New(s.config)))
	s.router.Handle("POST", "/admin/user/new", usershandlers.NewUser(store.New(s.config)))
	s.router.Handle("GET", "/admin/users/id", usershandlers.GetUserByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/users/delete/", usershandlers.DeleteUser(store.New(s.config)))

	s.router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels/id:id", hotelhandlers.GetHotelByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/pets", pethandlers.AllPetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets/id:id", pethandlers.GetPetByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms/id:id", roomhandlers.GetRoomByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats/id:id", seathandlers.GetSeatByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/bookings", bookinghandlers.AllBookingsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings/id:id", bookinghandlers.GetBookingByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/employees", employeehandlers.AllEmployeeHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees/id:id", employeehandlers.GetEmployeeByID(store.New(s.config)))

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))
}
