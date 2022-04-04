package store_test

import (
	"admin/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host     = "localhost"
	dbName   = "test_admin"
	user     = "user"
	password = "userpass"
	port     = "6543"
	sslMode  = "disable"
)

func TestStore_Open(t *testing.T) {
	s, _ := store.TestStore(t, host, dbName, user, password, port, sslMode)
	err := s.Open()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
