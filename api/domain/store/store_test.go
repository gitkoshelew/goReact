package store_test

import (
	"os"
	"testing"
)

var (
	host, port, dbName, user, password, sslMode string
)

func TestMain(m *testing.M) {

	host = os.Getenv("TEST_POSTGRES_HOST")
	dbName = os.Getenv("TEST_POSTGRES_DB")
	user = os.Getenv("TEST_POSTGRES_USER")
	password = os.Getenv("TEST_POSTGRES_PASSWORD")
	port = os.Getenv("TEST_POSTGRES_PORT")
	sslMode = os.Getenv("TEST_POSTGRES_SSLMODE")
}
