package webapp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Server struct {
		Host     string
		Port     int
		LogLevel string
		Store    string
	}
	DbConnection struct {
		Host     string
		Port     int
		Username string
		Password string
		DbName   string
		Sslmode  string
	}
}

// NewConfig get values from env
func (c *Config) NewConfig() {
	c.Server.Host = os.Getenv("ADMIN_SERVER_HOST")
	c.Server.Port, _ = strconv.Atoi(os.Getenv("ADMIN_SERVER_PORT"))
	c.Server.LogLevel = os.Getenv("ADMIN_SEVER_LOG_LEVEL")
	c.Server.Store = os.Getenv("ADMIN_SERVER_STORE")

	c.DbConnection.Host = os.Getenv("ADMIN_POSTGRES_HOST")
	c.DbConnection.DbName = os.Getenv("ADMIN_POSTGRES_DB")
	c.DbConnection.Username = os.Getenv("ADMIN_POSTGRES_USER")
	c.DbConnection.Password = os.Getenv("ADMIN_POSTGRES_PASSWORD")
	c.DbConnection.Port, _ = strconv.Atoi(os.Getenv("ADMIN_POSTGRES_PORT"))
	c.DbConnection.Sslmode = os.Getenv("ADMIN_POSTGRES_SSLMODE")

}

func init() {
	godotenv.Load("../../.env") // the path is true if you start application from current directory
}

// PgDataSource ...
func (c *Config) PgDataSource() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DbConnection.Host,
		c.DbConnection.Port,
		c.DbConnection.Username,
		c.DbConnection.Password,
		c.DbConnection.DbName,
		c.DbConnection.Sslmode,
	)
}

// ServerAddress ...
func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// ServerInfo ...
func (c *Config) ServerInfo() string {
	return fmt.Sprintf("Server Address: %s\nPG Data Source: %#v", c.ServerAddress(), c.PgDataSource())
}
