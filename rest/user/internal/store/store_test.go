package store_test

import (
	"testing"
	"user/internal/store"

	"github.com/stretchr/testify/assert"
)

var (
	host     = "localhost"
	dbName   = "userdbtest"
	user     = "user"
	password = "userpass"
	port     = "3231"
	sslMode  = "disable"
)

func TestStore_Open(t *testing.T) {
	s, _ := store.TestStore(t, host, dbName, user, password, port, sslMode)
	err := s.Open()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
