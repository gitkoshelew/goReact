package server

import (
	handlersadmin "goReact/webapp/admin/handlersAdmin"
	"net/http"
	//"goReact/webapp/server"
)

func (s *Server) configureRoutesAdmin() {
	s.router.HandlerFunc("GET", "/admin/login", handlersadmin.LoginAdmin())

	s.router.HandlerFunc("GET", "/admin/home", handlersadmin.HomeAdmin())

	//s.router.HandlerFunc("GET", "/admin/allusers", handlersadmin.LoginAdmin())
	
	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))

}
