package store_test

import (
	"hotel/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host                = "localhost"
	dbName              = "hoteldbtest"
	user                = "user"
	password            = "userpass"
	port                = "8088"
	sslMode             = "disable"
	testStore, teardown = store.TestStore(&testing.T{}, host, dbName, user, password, port, sslMode)
)

func TestStore_Open(t *testing.T) {
	s, _ := store.TestStore(t, host, dbName, user, password, port, sslMode)
	err := s.Open()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
