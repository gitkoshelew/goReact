package server

import (
	"goReact/webapp/server/handlers"
	"goReact/webapp/server/handlers/account"
	"goReact/webapp/server/handlers/booking"
	"goReact/webapp/server/handlers/client"
	"goReact/webapp/server/handlers/employee"
	"goReact/webapp/server/handlers/hotel"
	"goReact/webapp/server/handlers/pet"
	"goReact/webapp/server/handlers/room"
	"goReact/webapp/server/handlers/seat"
	"goReact/webapp/server/handlers/user"
)

func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/", handlers.HandleHomePage())

	s.router.HandlerFunc("GET", "/api/accounts", account.GetAccountsHandler())
	s.router.HandlerFunc("POST", "/api/account", account.PostAccountHandler())
	s.router.HandlerFunc("PUT", "/api/account", account.PutAccountHandler())
	s.router.Handle("GET", "/api/account/:id", account.GetAccountHandle())
	s.router.Handle("DELETE", "/api/account/:id", account.DeleteAccountHandle())

	s.router.HandlerFunc("GET", "/api/users", user.GetUsersHandler())
	s.router.HandlerFunc("POST", "/api/user", user.PostUserHandler())
	s.router.HandlerFunc("PUT", "/api/user", user.PutUserHandle())
	s.router.Handle("GET", "/api/user/:id", user.GetUserHandle())
	s.router.Handle("DELETE", "/api/user/:id", user.DeleteUserHandle())

	s.router.HandlerFunc("GET", "/api/employees", employee.GetEmployeesHandler())
	s.router.HandlerFunc("POST", "/api/employee", employee.PostEmployeesHandler())
	s.router.HandlerFunc("PUT", "/api/employee", employee.PutEmployeesHandler())
	s.router.Handle("GET", "/api/employee/:id", employee.GetEmployeeHandle())
	s.router.Handle("DELETE", "/api/employee/:id", employee.DeleteEmployeeHandle())

	s.router.HandlerFunc("GET", "/api/clients", client.GetClientsHandler())
	s.router.HandlerFunc("POST", "/api/client", client.PostClientsHandler())
	s.router.HandlerFunc("PUT", "/api/client", client.PutClientsHandler())
	s.router.Handle("GET", "/api/client/:id", client.GetClientHandle())
	s.router.Handle("DELETE", "/api/client/:id", client.DeleteClientHandle())

	s.router.HandlerFunc("GET", "/api/pets", pet.GetPetsHandler())
	s.router.HandlerFunc("POST", "/api/pet", pet.PostPetHandler())
	s.router.HandlerFunc("PUT", "/api/pet", pet.PutPetHandler())
	s.router.Handle("GET", "/api/pet/:id", pet.GetPetHandle())
	s.router.Handle("DELETE", "/api/pet/:id", pet.DeletePetHandle())

	s.router.HandlerFunc("GET", "/api/hotels", hotel.GetHotelsHandler())
	s.router.HandlerFunc("POST", "/api/hotel", hotel.PostHotelHandler())
	s.router.HandlerFunc("PUT", "/api/hotel", hotel.PutHotelHandler())
	s.router.Handle("GET", "/api/hotel/:id", hotel.GetHotelHandle())
	s.router.Handle("DELETE", "/api/hotel/:id", hotel.DeleteHotel())

	s.router.HandlerFunc("GET", "/api/rooms", room.GetRoomsHandler())
	s.router.HandlerFunc("POST", "/api/room", room.PostRoomHandler())
	s.router.HandlerFunc("PUT", "/api/room", room.PutRoomHandler())
	s.router.Handle("GET", "/api/room/:id", room.GetRoomHandle())
	s.router.Handle("DELETE", "/api/room/:id", room.DeleteRoomHandle())

	s.router.HandlerFunc("GET", "/api/seats", seat.GetSeatsHandler())
	s.router.HandlerFunc("POST", "/api/seat", seat.PostSeatHandler())
	s.router.HandlerFunc("PUT", "/api/seat", seat.PutSeatHandler())
	s.router.Handle("GET", "/api/seat/:id", seat.GetSeatHandle())
	s.router.Handle("DELETE", "/api/seat/:id", seat.DeleteSeatHandle())

	s.router.HandlerFunc("GET", "/api/bookings", booking.GetBookingsHandler())
	s.router.HandlerFunc("POST", "/api/booking", booking.PostBookingsHandler())
	s.router.HandlerFunc("PUT", "/api/booking", booking.PutBookingsHandler())
	s.router.Handle("GET", "/api/booking/:id", booking.GetBookingHandle())
	s.router.Handle("DELETE", "/api/booking/:id", booking.DeleteBookingHandle())
}
