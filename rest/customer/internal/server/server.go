package server

import (
	"customer/internal/store"
	"net/http"
	"customer/internal/config"
	"customer/pkg/logging"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	Config *config.Config
	Router *httprouter.Router
	Logger *logging.Logger
	Store  *store.Store
}

// New ...
func New(confing *config.Config) *Server {
	return &Server{
		Config: confing,
		Logger: logging.GetLogger(),
		Router: httprouter.New(),
		Store:  store.New(confing),
	}
}

// Start ...
func (s *Server) Start() error {
	s.ConfigureRouter()
	s.Logger.Info("Router started successfully")

	if err := s.Store.Open(); err != nil {
		s.Logger.Errorf("err during db opening, err: %w", err)
		return err
	}

	CORS := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:1110", "http://0.0.0.0:1110"},
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

	s.Logger.Infof("Auth server starts at %s ...", s.Config.Server.Address)
	handler := CORS.Handler(s.Router)
	return http.ListenAndServe(s.Config.Server.Address, handler)
}
