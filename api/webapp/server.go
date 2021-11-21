package webapp

import (
	"goReact/handler"
	"log"
	"net/http"
)

func StartServer(config *Config) {
	addr := config.ServerAddress()

	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler.New()))
}
