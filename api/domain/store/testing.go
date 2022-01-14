package store

import (
	"fmt"
	"goReact/webapp"
	"strconv"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, host, dbName, user, password, port, sslMode string) (*Store, func(...string)) {
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

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.Db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}

}
