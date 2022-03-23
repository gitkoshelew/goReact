package store

import (
	"booking/domain/model"
	"booking/internal/config"
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// ID ...
type ID struct {
	Booking int
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
	s.Logger.Infof("Booking store opening. Source: %s", dataSourceName)

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
		_, err := s.Db.Exec("TRUNCATE booking CASCADE")
		if err != nil {
			t.Fatal(err)
		}
		s.Close()
	}
}

// FillDB ...
func FillDB(t *testing.T, s *Store) *ID {
	s.Open()

	booking := model.TestBooking()
	bookingID, _ := s.Booking().Create(booking)
	booking.BookingID = *bookingID

	s.Close()

	return &ID{
		Booking: *bookingID,
	}
}
