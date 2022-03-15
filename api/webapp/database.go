package webapp

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ...
)

// DataBase ...
type DataBase struct {
	dataBase *sql.DB
	config   *Config
}

// NewDataBase ...
func NewDataBase(config *Config) *DataBase {
	dataBase, err := ConnectDb(config)
	if err != nil {
		log.Panic(err)
	}

	return &DataBase{
		dataBase: dataBase,
		config:   config,
	}
}

// ConnectDb ...
func ConnectDb(config *Config) (*sql.DB, error) {
	dataSourceName := config.PgDataSource()

	log.Printf("Connecting to database via %#v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Database connection successfull!")
	return db, err
}
