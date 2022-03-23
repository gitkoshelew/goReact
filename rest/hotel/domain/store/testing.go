package store

import (
	"database/sql"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/config"
	"os"
	"testing"
)

// ID ...
type ID struct {
	Employee int
	Hotel    int
	Room     int
	Seat     int
}

// TestStore ...
func TestStore(t *testing.T, host, dbName, user, password, port, sslMode string) (*Store, func()) {
	t.Helper()

	config := &config.Config{
		Server: struct{ Address string }{
			Address: fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVER_HOST"), os.Getenv("AUTH_SERVER_PORT")),
		},
		DataBase: struct {
			Host     string
			Port     string
			Username string
			Password string
			DbName   string
			Sslmode  string
		}{
			Host:     host,
			Port:     port,
			Username: user,
			Password: password,
			DbName:   dbName,
			Sslmode:  sslMode,
		}}

	s := New(config)

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.Config.DataBase.Host,
		s.Config.DataBase.Port,
		s.Config.DataBase.Username,
		s.Config.DataBase.Password,
		s.Config.DataBase.DbName,
		s.Config.DataBase.Sslmode,
	)
	s.Logger.Infof("Hotel store opening. Source: %s", dataSourceName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		s.Logger.Errorf("Can't open DB. Err msg: %v", err)
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		s.Logger.Errorf("Can't ping DB. Err msg: %v", err)
		t.Fatal(err)
	}

	return s, func() {
		s.Open()
		_, err := s.Db.Exec("TRUNCATE employee, hotel, room, seat CASCADE")
		if err != nil {
			t.Fatal(err)
		}
		s.Close()
	}
}

// FillDB ...
func FillDB(t *testing.T, s *Store) *ID {
	s.Open()

	hotel := model.TestHotel()
	hotelID, _ := s.Hotel().Create(hotel)
	hotel.HotelID = *hotelID

	employee := model.TestEmployee()
	employee.Hotel = *hotel
	employeeID, _ := s.Employee().Create(employee)
	employee.EmployeeID = *employeeID

	room := model.TestRoom()
	room.Hotel = *hotel
	roomID, _ := s.Room().Create(room)
	room.RoomID = *roomID

	seat := model.TestSeat()
	seat.Room = *room
	seatID, _ := s.Seat().Create(seat)
	seat.SeatID = *seatID

	s.Close()

	return &ID{
		Employee: *employeeID,
		Hotel:    *hotelID,
		Room:     *roomID,
		Seat:     *seatID,
	}
}
