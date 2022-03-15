package main

import (
	"hotel/internal/config"
	"hotel/internal/server"
)

func main() {

	config := config.Get()
	server := server.New(config)

	if err := server.Start(); err != nil {
		server.Logger.Errorf("failed to start server due to error: %w", err)
	}
}
