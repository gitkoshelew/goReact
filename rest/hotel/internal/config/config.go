package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Server struct {
		Address string
	}
	DataBase struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
		Sslmode  string
	}
}

func init() {
	godotenv.Load("../../../.env") // the path is true if you start application from current directory
}

// Get config
func Get() *Config {
	return &Config{
		Server: struct{ Address string }{
			Address: fmt.Sprintf("%s:%s", os.Getenv("HOTEL_SERVER_HOST"), os.Getenv("HOTEL_SERVER_PORT")),
		},
		DataBase: struct {
			Host     string
			Port     string
			Username string
			Password string
			DbName   string
			Sslmode  string
		}{
			Host:     os.Getenv("HOTEL_POSTGRES_HOST"),
			Port:     os.Getenv("HOTEL_POSTGRES_PORT"),
			Username: os.Getenv("HOTEL_POSTGRES_USER"),
			Password: os.Getenv("HOTEL_POSTGRES_PASSWORD"),
			DbName:   os.Getenv("HOTEL_POSTGRES_DB"),
			Sslmode:  os.Getenv("HOTEL_POSTGRES_SSLMODE"),
		},
	}
}
