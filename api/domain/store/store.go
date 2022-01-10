package store

import (
	"database/sql"
	"goReact/webapp"
	"goReact/webapp/server/logging"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	Config             *webapp.Config
	Db                 *sql.DB
	UserRepository     *UserRepository
	EmployeeRepository *EmployeeRepository
	HotelRepository    *HotelRepository
	RoomRepository     *RoomRepository
	SeatRepository     *SeatRepository
	PetRepository      *PetRepository
	BookingRepository  *BookingRepository
	ImageRepository    *ImageRepository
	Logger             *logging.Logger
}

// New ...
func New(config *webapp.Config) *Store {
	return &Store{
		Config: config,
		Logger: logging.GetLogger(),
	}
}

// Open ...
func (s *Store) Open() error {
	dataSourceName := s.Config.PgDataSource()
	s.Logger.Infof("Connecting to database via %#v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		s.Logger.Errorf("Error while trying open Data base. Msg: %v", err)
		return err
	}

	if err := db.Ping(); err != nil {
		s.Logger.Errorf("Error while ping to data base. MSG: %v", err)
		return err
	}
	s.Logger.Infof("Database connection successfull!")

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

// Booking ...
func (s *Store) Booking() *BookingRepository {
	if s.BookingRepository != nil {
		return s.BookingRepository
	}
	s.BookingRepository = &BookingRepository{
		Store: s,
	}
	return s.BookingRepository
}

// Image ...
func (s *Store) Image() *ImageRepository {
	if s.ImageRepository != nil {
		return s.ImageRepository
	}
	s.ImageRepository = &ImageRepository{
		Store: s,
	}
	return s.ImageRepository
}
