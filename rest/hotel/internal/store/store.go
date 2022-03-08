package store

import (
	"database/sql"
	"fmt"
	"hotel/internal/config"
	"hotel/pkg/logging"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	Config             *config.Config
	Db                 *sql.DB
	Logger             *logging.Logger
	HotelRepository    *HotelRepository
	RoomRepository     *RoomRepository
	SeatRepository     *SeatRepository
	EmployeeRepository *EmployeeRepository
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

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		return err
	}

	if err := db.Ping(); err != nil {
		s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		return err
	}

	s.Db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.Db.Close()
}

// Hotel ...
func (s *Store) Hotel() *HotelRepository {
	if s.HotelRepository != nil {
		return s.HotelRepository
	}
	s.HotelRepository = &HotelRepository{
		Store: s,
	}
	return s.HotelRepository
}

// Room ...
func (s *Store) Room() *RoomRepository {
	if s.RoomRepository != nil {
		return s.RoomRepository
	}
	s.RoomRepository = &RoomRepository{
		Store: s,
	}
	return s.RoomRepository
}

// Seat ...
func (s *Store) Seat() *SeatRepository {
	if s.SeatRepository != nil {
		return s.SeatRepository
	}
	s.SeatRepository = &SeatRepository{
		Store: s,
	}
	return s.SeatRepository
}

// Employee ...
func (s *Store) Employee() *EmployeeRepository {
	if s.EmployeeRepository != nil {
		return s.EmployeeRepository
	}

	s.EmployeeRepository = &EmployeeRepository{
		Store: s,
	}

	return s.EmployeeRepository
}
