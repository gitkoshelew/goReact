package server

import (
	"admin/domain/store"
	"admin/webapp"
	"admin/webapp/logger"
	"admin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	config *webapp.Config
	logger *logger.Logger
	router *httprouter.Router
	Store  *store.Store
}

// New ...
func New(config *webapp.Config) *Server {
	return &Server{
		config: config,
		logger: logger.GetLogger(),
		router: httprouter.New(),
	}
}

// Start ...
func (s *Server) Start() error {

	s.configureRoutes()
	s.logger.Info("Admin router started successfully")

	err := session.OpenSessionStore(s.config)
	if err != nil {
		s.logger.Errorf("Error while open sessions store. ERR MSG: %s", err.Error())
		return err
	}
	s.logger.Info("Session store started successfully")

	if err := s.configureStore(); err != nil {
		s.logger.Errorf("Error while configure store. ERR MSG: %s", err.Error())
		return err
	}
	s.logger.Info("Store started successfully")

	s.logger.Infof("Server starts at %s ...", s.config.ServerInfo())

	return http.ListenAndServe(s.config.ServerAddress(), s.router)
}
