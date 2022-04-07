package server

import (
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.router.Handle("POST", "/save", image.SaveJPEGHandle(store.New(s.config)))
	s.router.Handle("GET", "/getImage", image.GetImageHandle(store.New(s.config)))

}
