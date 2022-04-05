package store

import (
	"auth/domain/model"
	"auth/internal/config"
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// ID ...
type ID struct {
	User int
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
	s.Logger.Infof("Auth store opening. Source: %s", dataSourceName)

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
		_, err := s.Db.Exec("TRUNCATE users CASCADE")
		if err != nil {
			t.Fatal(err)
		}
		s.Close()
	}
}

// FillDB ...
func FillDB(t *testing.T, s *Store) *ID {
	s.Open()

	user := model.TestUser()
	user.Email = "test@mail.org"
	userID, _ := s.User().Create(user)
	user.UserID = *userID

	s.Close()

	return &ID{
		User: *userID,
	}
}
