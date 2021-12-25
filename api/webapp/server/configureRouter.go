package server

import (
	"goReact/webapp/server/handlers"
	"goReact/webapp/server/handlers/account"
	"goReact/webapp/server/handlers/authentication"
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
	s.router.Handle("GET", "/", handlers.HandleHomePage())
	s.router.Handle("POST", "/login", authentication.LoginHandle())
	s.router.Handle("POST", "/logout", authentication.LogoutHandle())
	s.router.Handle("POST", "/todo", authentication.TodoHandle())
	s.router.Handle("POST", "/refresh", authentication.RefreshHandle())

	s.router.Handle("GET", "/api/accounts", account.GetAccountsHandle())
	s.router.Handle("POST", "/api/account", account.PostAccountHandle())
	s.router.Handle("PUT", "/api/account", account.PutAccountHandle())
	s.router.Handle("GET", "/api/account/:id", account.GetAccountHandle())
	s.router.Handle("DELETE", "/api/account/:id", account.DeleteAccountHandle())

	s.router.Handle("GET", "/api/users", user.GetUsersHandle())
	s.router.Handle("POST", "/api/user", user.PostUserHandle())
	s.router.Handle("PUT", "/api/user", user.PutUserHandle())
	s.router.Handle("GET", "/api/user/:id", user.GetUserHandle())
	s.router.Handle("DELETE", "/api/user/:id", user.DeleteUserHandle())

	s.router.Handle("GET", "/api/employees", employee.GetEmployeesHandle())
	s.router.Handle("POST", "/api/employee", employee.PostEmployeesHandle())
	s.router.Handle("PUT", "/api/employee", employee.PutEmployeesHandle())
	s.router.Handle("GET", "/api/employee/:id", employee.GetEmployeeHandle())
	s.router.Handle("DELETE", "/api/employee/:id", employee.DeleteEmployeeHandle())

	s.router.Handle("GET", "/api/clients", client.GetClientsHandle())
	s.router.Handle("POST", "/api/client", client.PostClientsHandle())
	s.router.Handle("PUT", "/api/client", client.PutClientsHandle())
	s.router.Handle("GET", "/api/client/:id", client.GetClientHandle())
	s.router.Handle("DELETE", "/api/client/:id", client.DeleteClientHandle())

	s.router.Handle("GET", "/api/pets", pet.GetPetsHandle())
	s.router.Handle("POST", "/api/pet", pet.PostPetHandle())
	s.router.Handle("PUT", "/api/pet", pet.PutPetHandle())
	s.router.Handle("GET", "/api/pet/:id", pet.GetPetHandle())
	s.router.Handle("DELETE", "/api/pet/:id", pet.DeletePetHandle())

	s.router.Handle("GET", "/api/hotels", hotel.GetHotelsHandle())
	s.router.Handle("POST", "/api/hotel", hotel.PostHotelHandle())
	s.router.Handle("PUT", "/api/hotel", hotel.PutHotelHandle())
	s.router.Handle("GET", "/api/hotel/:id", hotel.GetHotelHandle())
	s.router.Handle("DELETE", "/api/hotel/:id", hotel.DeleteHotel())

	s.router.Handle("GET", "/api/rooms", room.GetRoomsHandle())
	s.router.Handle("POST", "/api/room", room.PostRoomHandle())
	s.router.Handle("PUT", "/api/room", room.PutRoomHandle())
	s.router.Handle("GET", "/api/room/:id", room.GetRoomHandle())
	s.router.Handle("DELETE", "/api/room/:id", room.DeleteRoomHandle())

	s.router.Handle("GET", "/api/seats", seat.GetSeatsHandle())
	s.router.Handle("POST", "/api/seat", seat.PostSeatHandle())
	s.router.Handle("PUT", "/api/seat", seat.PutSeatHandle())
	s.router.Handle("GET", "/api/seat/:id", seat.GetSeatHandle())
	s.router.Handle("DELETE", "/api/seat/:id", seat.DeleteSeatHandle())

	s.router.Handle("GET", "/api/bookings", booking.GetBookingsHandle())
	s.router.Handle("POST", "/api/booking", booking.PostBookingsHandle())
	s.router.Handle("PUT", "/api/booking", booking.PutBookingsHandle())
	s.router.Handle("GET", "/api/booking/:id", booking.GetBookingHandle())
	s.router.Handle("DELETE", "/api/booking/:id", booking.DeleteBookingHandle())
}
