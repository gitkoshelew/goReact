package server

import (
	"image/internal/handlers/image"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("GET", "/image", image.GetImageHandle(s.Store))
	s.Router.Handle("GET", "/images", image.GetAllImagesHandle(s.Store))
	s.Router.Handle("POST", "/image", image.SaveJPEGHandle(s.Store))
	s.Router.Handle("DELETE", "/image/:id", image.DeleteImageHandle(s.Store))
	s.Router.Handle("PUT", "/image", image.UpdateImage(s.Store))

}
