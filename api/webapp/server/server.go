package server

import (
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp"
	"goReact/webapp/server/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	config *webapp.Config
	Logger *logging.Logger
	router *httprouter.Router
	Store  *store.Store
	Mail   *service.Mail
}

// New ...
func New(config *webapp.Config) *Server {
	return &Server{
		config: config,
		Logger: logging.GetLogger(),
		router: httprouter.New(),
		Mail:   service.GetMail(config),
	}
}

// Start ...
func (s *Server) Start() error {

	s.configureRouter()
	s.Logger.Info("Router started successfully")

	s.configureRoutesAdmin()
	s.Logger.Info("Admin router started successfully")

	if err := s.configureStore(); err != nil {
		s.Logger.Errorf("Error while configure store. ERR MSG: %s", err.Error())
		return err
	}
	s.Logger.Info("Store started successfully")

	s.Logger.Infof("Server starts at %s ...", s.config.ServerInfo())
	CORS := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://react-ngnix-app:3001"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Access-Token"},
	})

	handler := CORS.Handler(s.router)
	return http.ListenAndServe(s.config.ServerAddress(), handler)
}
