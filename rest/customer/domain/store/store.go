package store

import (
	"customer/internal/config"
	"customer/pkg/logging"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // ...
)

var (
	// ErrEmailIsUsed ...
	ErrEmailIsUsed = errors.New("Email already in use")

	// ErrNoRowsAffected ...
	ErrNoRowsAffected = errors.New("No rows affected")
)

// Store ...
type Store struct {
	Config         *config.Config
	Db             *sql.DB
	Logger         *logging.Logger
	UserRepository *UserRepository
	PetRepository  *PetRepository
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
	s.Logger.Infof("User store opening. Source: %s", dataSourceName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
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

// Pet ...
func (s *Store) Pet() *PetRepository {
	if s.PetRepository != nil {
		return s.PetRepository
	}

	s.PetRepository = &PetRepository{
		Store: s,
	}

	return s.PetRepository
}
