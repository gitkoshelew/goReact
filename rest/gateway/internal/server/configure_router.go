package server

import (
	"gateway/internal/client"
	"gateway/internal/handlers/auth"
	"gateway/internal/handlers/booking"
	"gateway/internal/handlers/customer"
	"gateway/internal/handlers/middleware"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	// Authentication service handlers
	s.Router.Handle("POST", "/login", middleware.IsLoggedIn(auth.LoginHandle(client.AuthLoginService)))
	s.Router.Handle("POST", "/logout", auth.LogoutHandle(client.AuthLogoutService))
	s.Router.Handle("POST", "/registration", auth.RegistrationHandle(client.AuthRegistrationService))
	s.Router.Handle("POST", "/refresh", auth.RefreshHandle(client.AuthRefreshService))

	// Customer service handlers
	s.Router.Handle("GET", "/users", customer.GetAllUsersHandle(client.CustomerGetAllUsersService))
	s.Router.Handle("GET", "/user/:id", customer.GetUserHandle(client.CustomerUserService))
	s.Router.Handle("DELETE", "/user/:id", customer.DeleteUserHandle(client.CustomerUserService))
	s.Router.Handle("PUT", "/user", customer.UpdateUserHandle(client.CustomerUserService))
	s.Router.Handle("POST", "/user", customer.CreateUserHandle(client.CustomerUserService))

	s.Router.Handle("GET", "/pets", customer.GetAllPetsHandle(client.CustomerGetAllPetsService))
	s.Router.Handle("GET", "/pet/:id", customer.GetPetHandle(client.CustomerPetService))
	s.Router.Handle("DELETE", "/pet/:id", customer.DeletePetHandle(client.CustomerPetService))
	s.Router.Handle("PUT", "/pet", customer.UpdatePetHandle(client.CustomerPetService))
	s.Router.Handle("POST", "/pet", customer.CreatePetHandle(client.CustomerPetService))

	// Booking service handlers
	s.Router.Handle("GET", "/bookings", booking.GetAllBookingsHandle(client.BookingGetAllUsersService))
	s.Router.Handle("GET", "/booking/:id", booking.GetBookingHandle(client.BookingUserService))
	s.Router.Handle("DELETE", "/booking/:id", booking.DeleteBookingHandle(client.BookingUserService))
	s.Router.Handle("PUT", "/booking", booking.UpdateBookingHandle(client.BookingUserService))
	s.Router.Handle("POST", "/booking", booking.CreateBookingHandle(client.BookingUserService))

}
