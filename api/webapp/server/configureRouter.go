package server

import (
	"goReact/webapp/server/handlers"
)

func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/", handlers.HandleHomePage())
	// s.router.HandlerFunc("*", "/api", s.handleApi())

	s.router.HandlerFunc("GET", "/api/accounts", handlers.HandleAccounts())
	s.router.HandlerFunc("GET", "/api/account", handlers.HandleAccountSearch())

	s.router.HandlerFunc("GET", "/api/accounts/json", handlers.HandleAccountsJson())

	s.router.HandlerFunc("GET", "/api/users", handlers.HandleUsers())
	s.router.HandlerFunc("GET", "/api/user", handlers.HandleUserSearch())

	s.router.HandlerFunc("GET", "/api/employees", handlers.HandleEmployees())
	s.router.HandlerFunc("GET", "/api/employee", handlers.HandleEmployeeSearch())

	s.router.HandlerFunc("GET", "/api/clients", handlers.HandleClients())
	s.router.HandlerFunc("GET", "/api/client", handlers.HandleClientSearch())

	s.router.HandlerFunc("GET", "/api/pets", handlers.HandlePets())
	s.router.HandlerFunc("GET", "/api/pet", handlers.HandlePetSearch())

	s.router.HandlerFunc("GET", "/api/hotels", handlers.HandleHotels())
	s.router.HandlerFunc("GET", "/api/hotel", handlers.HandleHotelSearch())

	s.router.HandlerFunc("GET", "/api/rooms", handlers.HandleHotelRooms())
	s.router.HandlerFunc("GET", "/api/room", handlers.HandleHotelRoomSearch())

	s.router.HandlerFunc("GET", "/api/seats", handlers.HandleHotelRoomSeats())
	s.router.HandlerFunc("GET", "/api/seat", handlers.HandleHotelRoomSeatSearch())

	s.router.HandlerFunc("GET", "/api/bookings", handlers.HandleBookings())
	s.router.HandlerFunc("GET", "/api/booking", handlers.HandleBookingSearch())
}
