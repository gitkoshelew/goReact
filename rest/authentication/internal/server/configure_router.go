package server

import (
	"auth/domain/store"
	"auth/internal/handlers"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("POST", "/login", handlers.LoginHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/logout", handlers.LogoutHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/refresh", handlers.RefreshHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/registration", handlers.RegistrationHandle(store.New(s.Config)))
}
