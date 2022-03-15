package session

import (
	"admin/webapp"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
)

// SessionStore ...
type SessionStore struct {
	DB      *sql.DB
	PGStore *pgstore.PGStore
}

var sstore SessionStore

// OpenSessionStore ...
func OpenSessionStore(c *webapp.Config) error {
	defer sstore.PGStore.StopCleanup(sstore.PGStore.Cleanup(time.Minute * 5))

	dataSourceName := c.PgDataSource()
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Printf("Error occurred while opening pg session store. Err msg:%v.", err)
		return err
	}
	if err := db.Ping(); err != nil {
		log.Printf("Error occurred while connecting to pg session store. Err msg:%v.", err)
		return err
	}
	sstore.DB = db

	PGStore, err := pgstore.NewPGStoreFromPool(db, []byte(os.Getenv("ADMIN_SESSION_KEY")))
	if err != nil {
		log.Printf("Error occurred while creating pg session store. Err msg:%v.", err)
		return err
	}
	sstore.PGStore = PGStore

	return nil

}
