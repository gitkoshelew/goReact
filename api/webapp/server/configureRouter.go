package server

import (
	"goReact/domain/store"
	"goReact/webapp/server/handler/authentication"
	"goReact/webapp/server/handler/booking"
	"goReact/webapp/server/handler/hotel"
	"goReact/webapp/server/handler/image"
	"goReact/webapp/server/handler/middleware"
	gitoauth2 "goReact/webapp/server/handler/middleware/gitoauth2"
	linkedinoauth2 "goReact/webapp/server/handler/middleware/linkedinoauth2"


	"goReact/webapp/server/handler/pet"
	restorepassword "goReact/webapp/server/handler/restorePassword"
	"goReact/webapp/server/handler/room"
	"goReact/webapp/server/handler/seat"
	"goReact/webapp/server/handler/user"
)

func (s *Server) configureRouter() {

	s.router.Handler("POST", "/api/login", middleware.IsLoggedIn(middleware.ValidateLogin(authentication.LoginHandle(store.New(s.config)), store.New(s.config))))

	s.router.Handle("POST", "/save", image.SaveJPEGHandle(store.New(s.config)))
	s.router.Handle("GET", "/getImage", image.GetImageHandle(store.New(s.config)))

	s.router.Handle("POST", "/api/registration", middleware.ValidateUser(authentication.RegistrationHandle(store.New(s.config), s.Mail), store.New(s.config)))
	s.router.Handle("POST", "/api/logout", authentication.LogoutHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/refresh", authentication.RefreshHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/me", authentication.MeHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/users", user.GetUsersHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/user", middleware.ValidateUser(user.PostUserHandle(store.New(s.config)), store.New(s.config)))
	s.router.Handle("GET", "/api/user/:id", user.GetUserHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/pets/", pet.GetPetsHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/seat/:id", seat.GetSeatHandle(store.New(s.config)))
	s.router.Handle("GET", "/api/seats", seat.GetAllSeatsHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/seats/search/free", middleware.ValidateFreeSeatsSearchingRequest(seat.GetFreeSeatsHandle(store.New(s.config)), store.New(s.config)))

	s.router.Handle("GET", "/api/hotel/:id", hotel.GetHotelByID(store.New(s.config)))
	s.router.Handle("GET", "/api/hotels", hotel.AllHotelsHandler(store.New(s.config)))

	s.router.Handle("GET", "/api/room/:id", room.GetRoomHandle(store.New(s.config)))
	s.router.Handle("GET", "/api/rooms", room.GetAllRoomsHandle(store.New(s.config)))
	//localhost:8080/api/rooms/?offset=2&pagesize=2
	s.router.Handle("GET", "/api/rooms/", room.GetRoomsHandlePagination(store.New(s.config)))
	s.router.Handle("GET", "/api/toprooms/", room.GetTopRoomsHandle(store.New(s.config)))

	s.router.Handle("POST", "/api/booking", middleware.ValidateBooking(booking.PostBookingHandle(store.New(s.config)), store.New(s.config)))
	s.router.Handle("GET", "/api/bookings", booking.GetAllBookingsHandle(store.New(s.config)))
	s.router.Handle("GET", "/api/booking/:id", booking.GetBookingByIDHandler(store.New(s.config)))

	s.router.Handle("GET", "/api/emailconfirm/:token", authentication.EmailConfirm(store.New(s.config)))
	s.router.Handle("POST", "/api/forgotpassword", restorepassword.ForgotPassword(store.New(s.config), s.Mail))
	s.router.Handle("GET", "/api/emailrestore/:token", restorepassword.СhekingLinkForRestorePassword(store.New(s.config), restorepassword.ChangePassword(store.New(s.config))))

	s.router.Handle("GET", "/api/gitlogin", gitoauth2.GitHubLogin(store.New(s.config)))
	s.router.Handle("GET", "/api/gitlogin/re", gitoauth2.GitHubAuth(gitoauth2.GetUserGit(middleware.CreateToken(store.New(s.config)),store.New(s.config)), store.New(s.config)))

	s.router.Handle("GET", "/api/linkedinlogin", linkedinoauth2.LinkedInLogin(store.New(s.config)))
	s.router.Handle("GET", "/api/linkedinlogin/re", linkedinoauth2.LinkedInAuth(linkedinoauth2.GetUserLinkedIn(middleware.CreateToken(store.New(s.config)),store.New(s.config)), store.New(s.config)))
}
