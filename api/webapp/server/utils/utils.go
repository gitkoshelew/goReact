package utils

import (
	"database/sql"
	"goReact/webapp"
	"log"
)

// HandlerDbConnection returns DB
func HandlerDbConnection() *sql.DB {
	config := &webapp.Config{}
	config.NewConfig()
	db, err := webapp.ConnectDb(config)
	if err != nil {
		log.Fatal("Connection to db failed: ", err.Error())
	}
	return db
}
