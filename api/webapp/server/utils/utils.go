package utils

import (
	"database/sql"
	"goReact/webapp"
	"log"
	"net/http"

	"github.com/go-redis/redis"
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

// RedisConnection returns DB
func RedisConnection() *redis.Client {
	config := &webapp.Config{}
	dsn := config.RedisInfo()

	var client *redis.Client

	client = redis.NewClient(&redis.Options{Addr: dsn})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf(err.Error())
	}
	return client
}

// EnableCors ...
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
