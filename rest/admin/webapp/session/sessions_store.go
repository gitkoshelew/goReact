package session

import (
	"admin/webapp"
	"database/sql"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
)

type SessionStore struct {
	DB      *sql.DB
	PGStore *pgstore.PGStore
}

var sstore SessionStore

func OpenSessionStore(c *webapp.Config) error {
	defer sstore.PGStore.StopCleanup(sstore.PGStore.Cleanup(time.Minute * 5))

	dataSourceName := c.PgDataSource()
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	sstore.DB = db

	PGStore, err := pgstore.NewPGStoreFromPool(db, []byte(os.Getenv("ADMIN_SESSION_KEY")))
	if err != nil {
		return err
	}
	sstore.PGStore = PGStore

	return nil

}


