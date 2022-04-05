package server

import (
	"auth/domain/store"
	"auth/internal/handlers"
	"auth/internal/middleware"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handler("POST", "/login", middleware.IsLoggedIn(middleware.ValidateLogin(handlers.LoginHandle(store.New(s.Config)), store.New(s.Config))))
	s.Router.Handle("POST", "/logout", handlers.LogoutHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/refresh", handlers.RefreshHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/registration", middleware.ValidateUser(handlers.RegistrationHandle(store.New(s.Config)), store.New(s.Config)))
}
