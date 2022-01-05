package server

import (
	handlersadmin "goReact/webapp/admin/handlersAdmin"
	hotelhandlers "goReact/webapp/admin/handlersAdmin/hotelHandlers"
	pethandlers "goReact/webapp/admin/handlersAdmin/petHandlers"
	roomhandlers "goReact/webapp/admin/handlersAdmin/roomHandlers"
	seathandlers "goReact/webapp/admin/handlersAdmin/seatHandlers"
	usershandlers "goReact/webapp/admin/handlersAdmin/usersHandlers"
	"net/http"
)

func (s *Server) configureRoutesAdmin() {
	s.router.Handle("GET", "/admin/login", handlersadmin.LoginAdmin())
	s.router.Handle("POST", "/admin/auth", handlersadmin.AuthAdmin())
	s.router.Handle("GET", "/admin/logout", handlersadmin.LogoutAdmin())
	

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

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))


	
}
