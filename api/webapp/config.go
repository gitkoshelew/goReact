package webapp

import (
	"fmt"
	"os"
	"strconv"
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
	}
}

// NewConfig ..
func (c *Config) NewConfig() {
	c.Server.Host = os.Getenv("SERVER_HOST")
	c.Server.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	c.Server.LogLevel = os.Getenv("SEVER_LOG_LEVEL")
	c.Server.Store = os.Getenv("SERVER_STORE")

	c.DbConnection.Host = os.Getenv("POSTGRES_HOST")
	c.DbConnection.DbName = os.Getenv("POSTGRES_DB")
	c.DbConnection.Username = os.Getenv("POSTGRES_USER")
	c.DbConnection.Password = os.Getenv("POSTGRES_PASSWORD")
	c.DbConnection.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

}

// PgDataSource ...
func (c *Config) PgDataSource() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DbConnection.Host,
		c.DbConnection.Port,
		c.DbConnection.Username,
		c.DbConnection.Password,
		c.DbConnection.DbName,
	)
}

// ServerAddress ...
func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c *Config) String() string {
	return fmt.Sprintf("Server Address: %s\nPG Data Source: %#v", c.ServerAddress(), c.PgDataSource())
}
