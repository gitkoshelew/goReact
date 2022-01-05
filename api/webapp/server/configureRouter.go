package server

import (
	"goReact/domain/store"
	"goReact/webapp/server/handlers"
	"goReact/webapp/server/handlers/authentication"
	"goReact/webapp/server/handlers/middleware"
	"goReact/webapp/server/handlers/user"
)

func (s *Server) configureRouter() {
	s.router.GlobalOPTIONS = handlers.CorsHandle()

	s.router.Handle("GET", "/", handlers.HandleHomePage())

	s.router.Handler("GET", "/private/whoami", middleware.Private(middleware.WhoAmI(), store.New(s.config)))
	s.router.Handler("POST", "api/login", middleware.IsLoggedIn(authentication.LoginHandle(store.New(s.config))))

	s.router.Handle("POST", "/api/registration", authentication.RegistrationHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/logout", authentication.LogoutHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/refresh", authentication.RefreshHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/todo", authentication.TodoHandle(store.New(s.config)))

	s.router.Handle("GET", "/api/users", user.GetUsersHandle(store.New(s.config)))
	s.router.Handle("POST", "/api/user", user.PostUserHandle(store.New(s.config)))
	s.router.Handle("GET", "/api/user/:id", user.GetUserHandle(store.New(s.config)))
}

