package store

import (
	"admin/webapp"
	"admin/webapp/logger"
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	Config                        *webapp.Config
	Db                            *sql.DB
	UserRepository                *UserRepository
	EmployeeRepository            *EmployeeRepository
	HotelRepository               *HotelRepository
	RoomRepository                *RoomRepository
	SeatRepository                *SeatRepository
	PetRepository                 *PetRepository
	BookingRepository             *BookingRepository
	ImageRepository               *ImageRepository
	PermissionsRepository         *PermissionsRepository
	PermissionsEmployeeRepository *PermissionsEmployeeRepository
	SessionRepository             *SessionRepository
	Logger                        *logger.Logger
}

// New ...
func New(config *webapp.Config) *Store {
	return &Store{
		Config: config,
		Logger: logger.GetLogger(),
	}
}

// Open ...
func (s *Store) Open() error {
	dataSourceName := s.Config.PgDataSource()
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

// Session ...
func (s *Store) Session() *SessionRepository {
	if s.SessionRepository != nil {
		return s.SessionRepository
	}

	s.SessionRepository = &SessionRepository{
		Store: s,
	}

	return s.SessionRepository
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

// Permissions ...
func (s *Store) Permissions() *PermissionsRepository {
	if s.PermissionsRepository != nil {
		return s.PermissionsRepository
	}
	s.PermissionsRepository = &PermissionsRepository{
		Store: s,
	}
	return s.PermissionsRepository
}

// PermissionsEmployee ...
func (s *Store) PermissionsEmployee() *PermissionsEmployeeRepository {
	if s.PermissionsEmployeeRepository != nil {
		return s.PermissionsEmployeeRepository
	}
	s.PermissionsEmployeeRepository = &PermissionsEmployeeRepository{
		Store: s,
	}
	return s.PermissionsEmployeeRepository
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
