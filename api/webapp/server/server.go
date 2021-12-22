package server

import (
	"goReact/webapp"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Server ...
type Server struct {
	config *webapp.Config
	logger *log.Logger
	router *httprouter.Router
}

// New ...
func New(config *webapp.Config) *Server {
	return &Server{
		config: config,
		logger: log.New(os.Stdout, "http: ", log.Ldate|log.Ltime|log.Lmsgprefix),
		router: httprouter.New(),
	}
}

// Start ...
func (s *Server) Start() error {

	s.configureRouter()

	s.logger.Printf("Server starting ...")

	return http.ListenAndServe(s.config.ServerAddress(), s.router)
}
