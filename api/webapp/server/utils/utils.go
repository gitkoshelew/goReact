package utils

import (
	"database/sql"
	"goReact/webapp"
	"log"
	"net/http"
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

// EnableCors ...
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
