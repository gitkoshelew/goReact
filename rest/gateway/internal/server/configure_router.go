package server

import (
	"gateway/internal/client"
	"gateway/internal/handlers/auth"
	"gateway/internal/handlers/booking"
	"gateway/internal/handlers/customer"
	"gateway/internal/handlers/hotel"
	"gateway/internal/handlers/middleware"
	"gateway/internal/handlers/image"
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

	// Hotel service handlers
	s.Router.Handle("GET", "/hotels", hotel.GetAllHotelsHandle(client.HotelGetAllHotelsService))
	s.Router.Handle("GET", "/hotel/:id", hotel.GetHotelHandle(client.HotelHotelService))
	s.Router.Handle("DELETE", "/hotel/:id", hotel.DeleteHotelHandle(client.HotelHotelService))
	s.Router.Handle("PUT", "/hotel", hotel.UpdateHotelHandle(client.HotelHotelService))
	s.Router.Handle("POST", "/hotel", hotel.CreateHotelHandle(client.HotelHotelService))

	s.Router.Handle("GET", "/rooms", hotel.GetAllRoomsHandle(client.HotelGetAllRoomsService))
	s.Router.Handle("GET", "/room/:id", hotel.GetRoomHandle(client.HotelRoomService))
	s.Router.Handle("DELETE", "/room/:id", hotel.DeleteRoomHandle(client.HotelRoomService))
	s.Router.Handle("PUT", "/room", hotel.UpdateRoomHandle(client.HotelRoomService))
	s.Router.Handle("POST", "/room", hotel.CreateRoomHandle(client.HotelRoomService))

	s.Router.Handle("GET", "/seats", hotel.GetAllSeatsHandle(client.HotelGetAllSeatsService))
	s.Router.Handle("GET", "/seat/:id", hotel.GetSeatHandle(client.HotelSeatService))
	s.Router.Handle("DELETE", "/seat/:id", hotel.DeleteSeatHandle(client.HotelSeatService))
	s.Router.Handle("PUT", "/seat", hotel.UpdateSeatHandle(client.HotelSeatService))
	s.Router.Handle("POST", "/seat", hotel.CreateSeatHandle(client.HotelSeatService))

	s.Router.Handle("GET", "/employees", hotel.GetAllEmployeesHandle(client.HotelGetAllEmployeesService))
	s.Router.Handle("GET", "/employee/:id", hotel.GetEmployeeHandle(client.HotelEmployeeService))
	s.Router.Handle("DELETE", "/employee/:id", hotel.DeleteEmployeeHandle(client.HotelEmployeeService))
	s.Router.Handle("PUT", "/employee", hotel.UpdateEmployeeHandle(client.HotelEmployeeService))
	s.Router.Handle("POST", "/employee", hotel.CreateEmployeeHandle(client.HotelEmployeeService))

	s.Router.Handle("GET", "/images", image.GetAllImagesHandle(client.GetAllImagesService))
	s.Router.Handle("GET", "/image/", image.GetImageHandle(client.ImageService))
	s.Router.Handle("DELETE", "/image/:id", image.DeleteImageHandle(client.ImageService))
	s.Router.Handle("PUT", "/image", image.UpdateimageHandle(client.ImageService))
	s.Router.Handle("POST", "/image", image.CreateImageHandle(client.ImageService))
}
