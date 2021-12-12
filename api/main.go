package main

import (
	"goReact/webapp"
	"goReact/webapp/server"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Need this part? If we have .env in docker-compose.yaml
	err := godotenv.Load(filepath.Join("../", ".env"))
	if err != nil {
		log.Print("Error loading .env file")
	}

	config := &webapp.Config{}
	config.NewConfig()
	webapp.ConnectDb(config)

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
