package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Server struct {
		Address string
	}

	ServiceGoreact struct {
		BaseURL             string
		ResourceGetAllUsers string
	}
}

func init() {
	err := godotenv.Load("../../../.env") // the path is true if you start application from current directory
	if err != nil {
		log.Printf("err: %v", err)
	}
}

// Get config
func Get() *Config {
	return &Config{
		Server: struct{ Address string }{
			Address: fmt.Sprintf("%s:%s", os.Getenv("GATEWAY_SERVER_HOST"), os.Getenv("GATEWAY_SERVER_PORT")),
		},
		ServiceGoreact: struct {
			BaseURL             string
			ResourceGetAllUsers string
		}{
			BaseURL:             os.Getenv("SERVICE_GOREACT_BASE_URL"),
			ResourceGetAllUsers: os.Getenv("SERVICE_GOREACT_RESOURCE_GETALLUSERS"),
		},
	}
}
