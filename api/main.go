package main

import (
	"goReact/webapp"
	"goReact/webapp/server"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(filepath.Join("../", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &webapp.Config{}
	config.NewConfig()

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
