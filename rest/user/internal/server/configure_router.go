package server

import (
	pethandlers "user/internal/handlers/petHandlers"
	usershandlers "user/internal/handlers/usersHandlers"
	"user/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/pets", pethandlers.AllPetsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/pet/:id", pethandlers.GetPetByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/pet/:id", pethandlers.DeletePet(store.New(s.Config)))
	s.Router.Handle("POST", "/pet", pethandlers.CreatePet(store.New(s.Config)))
	s.Router.Handle("PUT", "/pet", pethandlers.UpdatePet(store.New(s.Config)))

	s.Router.Handle("GET", "/users", usershandlers.AllUsersHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/user/:id", usershandlers.GetUserByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/user/:id", usershandlers.DeleteUser(store.New(s.Config)))
	s.Router.Handle("PUT", "/user", usershandlers.UpdateUser(store.New(s.Config)))
	s.Router.Handle("POST", "/user", usershandlers.NewUser(store.New(s.Config)))

}
