package main

import (
	"goReact/webapp"
	"goReact/webapp/server"
	"log"
)

func main() {

	config := &webapp.Config{}
	config.ReadFromFile("config.yaml")

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
