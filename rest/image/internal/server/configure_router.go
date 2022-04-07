package server

import (
	"image/internal/handlers/image"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("GET", "image/image", image.GetImageHandle(s.Store))
	s.Router.Handle("GET", "image/images", image.GetAllImagesHandle(s.Store))
	s.Router.Handle("POST", "image/save", image.SaveJPEGHandle(s.Store))
	s.Router.Handle("DELETE", "/image/:id", image.DeleteImageHandle(s.Store))
	s.Router.Handle("PUT", "/image", image.DeleteImageHandle(s.Store))

}

/*	s.Router.Handle("GET", "/employees", employeehandlers.AllEmployeeHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/employee/:id", employeehandlers.GetEmployeeByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/employee/:id", employeehandlers.DeleteEmployee(store.New(s.Config)))
	s.Router.Handle("POST", "/employee", employeehandlers.CreateEmployee(store.New(s.Config)))
	s.Router.Handle("PUT", "/employee", employeehandlers.UpdateEmployee(store.New(s.Config)))*/
