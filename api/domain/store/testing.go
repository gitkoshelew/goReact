package store

import (
	"goReact/webapp"
	"strconv"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, host, dbName, user, password, port, sslMode string) *Store {
	t.Helper()

	config := &webapp.Config{}
	config.NewConfig()
	config.DbConnection.Host = host
	config.DbConnection.DbName = dbName
	config.DbConnection.Username = user
	config.DbConnection.Password = password
	config.DbConnection.Port, _ = strconv.Atoi(port)
	config.DbConnection.Sslmode = sslMode

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s
}
