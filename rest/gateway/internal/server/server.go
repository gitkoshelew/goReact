package server

import (
	"gateway/internal/config"
	"gateway/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	Config *config.Config
	Router *httprouter.Router
	Logger *logging.Logger
}

// New ...
func New(confing *config.Config) *Server {
	return &Server{
		Config: confing,
		Router: httprouter.New(),
		Logger: logging.GetLogger(),
	}
}

// Start ...
func (s *Server) Start() error {
	s.ConfigureRouter()
	s.Logger.Info("Router started successfully")

	CORS := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:1111"},
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

	handler := CORS.Handler(s.Router)
	s.Logger.Infof("Gateway server starts at %s ...", s.Config.Server.Address)
	return http.ListenAndServe(s.Config.Server.Address, handler)
}
