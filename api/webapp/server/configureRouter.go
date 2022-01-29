package server

import (
	"goReact/domain/store"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/authentication"
	"goReact/webapp/server/handler/authentication/restorePassword"
	"goReact/webapp/server/handler/middleware"
	"goReact/webapp/server/handler/pet"
	"goReact/webapp/server/handler/room"
	"goReact/webapp/server/handler/user"
)

func (s *Server) configureRouter() {

	s.router.Handle("GET", "/", handler.HandleHomePage())

	s.router.Handler("POST", "/api/login", middleware.IsLoggedIn(authentication.LoginHandle(store.New(s.config))))

	s.router.Handle("POST", "/api/registration", authentication.RegistrationHandle(store.New(s.config), s.Mail))
	s.router.Handle("POST", "/api/logout", authentication.LogoutHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/refresh", authentication.RefreshHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/me", authentication.MeHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/users", user.GetUsersHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/user", user.PostUserHandle(store.New(s.config)))
	s.router.Handle("GET", "/api/user/:id", user.GetUserHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/pets/", pet.GetPetsHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/rooms", room.GetAllRoomsHandle(store.New(s.config)))
	//localhost:8080/api/rooms/?offset=2&pagesize=2
	s.router.Handle("GET", "/api/rooms/", room.GetRoomsHandlePagination(store.New(s.config)))

	s.router.Handle("GET", "/api/emailconfirm/:token", authentication.EmailConfirm(store.New(s.config)))
	s.router.Handle("POST", "/api/forgotpassword", restorePassword.ForgotPassword(store.New(s.config), s.Mail))
	s.router.Handle("GET", "/api/emailrestore/:token", restorePassword.СhekingLinkForRestorePassword(store.New(s.config), restorePassword.ChangePassword(store.New(s.config))))
}

