package store

import (
	"auth/internal/config"
	"auth/pkg/logging"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	Config         *config.Config
	Db             *sql.DB
	Logger         *logging.Logger
	UserRepository *UserRepository
}

// New ...
func New(config *config.Config) *Store {
	return &Store{
		Config: config,
		Logger: logging.GetLogger(),
	}
}

// Open ...
func (s *Store) Open() error {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.Config.DataBase.Host,
		s.Config.DataBase.Port,
		s.Config.DataBase.Username,
		s.Config.DataBase.Password,
		s.Config.DataBase.DbName,
		s.Config.DataBase.Sslmode,
	)
	s.Logger.Infof("Auth store opening. Source: %s", dataSourceName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		s.Logger.Errorf("Can't open DB. Err msg: %w", err)
		return err
	}

	if err := db.Ping(); err != nil {
		s.Logger.Errorf("Can't ping DB. Err msg: %w", err)
		return err
	}

	s.Db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.Db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		Store: s,
	}

	return s.UserRepository
}
