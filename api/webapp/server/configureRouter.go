package server

import (
	"goReact/webapp/server/handlers"
)

func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/", handlers.HandleHomePage())

	s.router.HandlerFunc("GET", "/api/accounts", handlers.HandleAccounts())
	s.router.HandlerFunc("POST", "/api/account", handlers.HandleAccounts())
	s.router.HandlerFunc("PUT", "/api/account", handlers.HandleAccounts())
	s.router.Handle("GET", "/api/account/:id", handlers.HandleAccount())
	s.router.Handle("DELETE", "/api/account/:id", handlers.HandleAccount())

	s.router.HandlerFunc("GET", "/api/users", handlers.HandleUsers())
	s.router.HandlerFunc("POST", "/api/user", handlers.HandleUsers())
	s.router.HandlerFunc("PUT", "/api/user", handlers.HandleUsers())
	s.router.Handle("GET", "/api/user/:id", handlers.HandleUser())
	s.router.Handle("DELETE", "/api/user/:id", handlers.HandleUser())

	s.router.HandlerFunc("GET", "/api/employees", handlers.HandleEmployees())
	s.router.HandlerFunc("POST", "/api/employee", handlers.HandleEmployees())
	s.router.HandlerFunc("PUT", "/api/employee", handlers.HandleEmployees())
	s.router.Handle("GET", "/api/employee/:id", handlers.HandleEmployee())
	s.router.Handle("DELETE", "/api/employee/:id", handlers.HandleEmployee())

	s.router.HandlerFunc("GET", "/api/clients", handlers.HandleClients())
	s.router.HandlerFunc("POST", "/api/client", handlers.HandleClients())
	s.router.HandlerFunc("PUT", "/api/client", handlers.HandleClients())
	s.router.Handle("GET", "/api/client/:id", handlers.HandleClient())
	s.router.Handle("DELETE", "/api/client/:id", handlers.HandleClient())

	s.router.HandlerFunc("GET", "/api/pets", handlers.HandlePets())
	s.router.HandlerFunc("POST", "/api/pet", handlers.HandlePets())
	s.router.HandlerFunc("PUT", "/api/pet", handlers.HandlePets())
	s.router.Handle("GET", "/api/pet/:id", handlers.HandlePet())
	s.router.Handle("DELETE", "/api/pet/:id", handlers.HandlePet())

	s.router.HandlerFunc("GET", "/api/hotels", handlers.HandleHotels())
	s.router.HandlerFunc("POST", "/api/hotel", handlers.HandleHotels())
	s.router.HandlerFunc("PUT", "/api/hotel", handlers.HandleHotels())
	s.router.Handle("GET", "/api/hotel/:id", handlers.HandleHotel())
	s.router.Handle("DELETE", "/api/hotel/:id", handlers.HandleHotel())

	s.router.HandlerFunc("GET", "/api/rooms", handlers.HandleHotelRooms())
	s.router.HandlerFunc("POST", "/api/room", handlers.HandleHotelRooms())
	s.router.HandlerFunc("PUT", "/api/room", handlers.HandleHotelRooms())
	s.router.Handle("GET", "/api/room/:id", handlers.HandleHotelRoom())
	s.router.Handle("DELETE", "/api/room/:id", handlers.HandleHotelRoom())

	s.router.HandlerFunc("GET", "/api/seats", handlers.HandleHotelRoomSeats())
	s.router.HandlerFunc("POST", "/api/seat", handlers.HandleHotelRoomSeats())
	s.router.HandlerFunc("PUT", "/api/seat", handlers.HandleHotelRoomSeats())
	s.router.Handle("GET", "/api/seat/:id", handlers.HandleHotelRoomSeat())
	s.router.Handle("DELETE", "/api/seat/:id", handlers.HandleHotelRoomSeat())

	s.router.HandlerFunc("GET", "/api/bookings", handlers.HandleBookings())
	s.router.HandlerFunc("POST", "/api/booking", handlers.HandleBookings())
	s.router.HandlerFunc("PUT", "/api/booking", handlers.HandleBookings())
	s.router.Handle("GET", "/api/booking/:id", handlers.HandleBooking())
	s.router.Handle("DELETE", "/api/booking/:id", handlers.HandleBooking())
}
