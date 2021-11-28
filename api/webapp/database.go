package webapp

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ...
)

// ConnectDb ...
func ConnectDb(config *Config) (*sql.DB, error) {
	dataSourceName := config.PgDataSource()

	log.Printf("Connecting to database via %#v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}
