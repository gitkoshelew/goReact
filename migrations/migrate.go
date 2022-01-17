package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // ...
	_ "github.com/golang-migrate/migrate/v4/source/file"       // ...
)

// Up migrations
func Up() {
	m, err := migrate.New(
		"file://migrations/",
		"postgres://user:userpass@postgresql_database:5432/goreact?sslmode=disable")
	if err != nil {
		log.Printf("Migrate error: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Printf("Migrate error: %v", err)
	}
	log.Printf("Migrate UP")
}
