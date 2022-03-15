package server

import "booking/internal/handlers"

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("GET", "/bookings", handlers.GetAllHandle(s.Store))
	s.Router.Handle("GET", "/booking/:id", handlers.GetByIDHandle(s.Store))
	s.Router.Handle("DELETE", "/booking/:id", handlers.DeleteHandle(s.Store))
	s.Router.Handle("PUT", "/booking", handlers.UpdateHandle(s.Store))
	s.Router.Handle("POST", "/booking", handlers.CreateHandle(s.Store))
}
