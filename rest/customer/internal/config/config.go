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
			Address: fmt.Sprintf("%s:%s", os.Getenv("CUSTOMER_SERVER_HOST"), os.Getenv("CUSTOMER_SERVER_PORT")),
		},
		DataBase: struct {
			Host     string
			Port     string
			Username string
			Password string
			DbName   string
			Sslmode  string
		}{
			Host:     os.Getenv("CUSTOMER_POSTGRES_HOST"),
			Port:     os.Getenv("CUSTOMER_POSTGRES_PORT"),
			Username: os.Getenv("CUSTOMER_POSTGRES_USER"),
			Password: os.Getenv("CUSTOMER_POSTGRES_PASSWORD"),
			DbName:   os.Getenv("CUSTOMER_POSTGRES_DB"),
			Sslmode:  os.Getenv("CUSTOMER_POSTGRES_SSLMODE"),
		},
	}
}
