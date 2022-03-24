package store

import (
	"hotel/internal/config"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, host, dbName, user, password, port, sslMode string) (*Store, func()) {
	t.Helper()

	config := config.Get()
	//config.init()
	config.DataBase.Host = host
	config.DataBase.DbName = dbName
	config.DataBase.Username = user
	config.DataBase.Password = password
	config.DataBase.Port = port
	config.DataBase.Sslmode = sslMode

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func() {
		_, err := s.Db.Exec("TRUNCATE users CASCADE")
		if err != nil {
			t.Fatal(err)
		}
		s.Close()
	}
}
