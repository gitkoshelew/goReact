package webapp

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Server struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		LogLevel string `yaml:"log_level"`
	} `yaml:"server"`
	PgConnection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
	} `yaml:"pgConnection"`
}

// PgDataSource ...
func (c *Config) PgDataSource() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PgConnection.Host,
		c.PgConnection.Port,
		c.PgConnection.Username,
		c.PgConnection.Password,
		c.PgConnection.DbName,
	)
}

// ServerAddress ...
func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c *Config) String() string {
	return fmt.Sprintf("Server Address: %s\nPG Data Source: %#v", c.ServerAddress(), c.PgDataSource())
}

// ReadFromFile ...
func (c *Config) ReadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open config file due to error: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Could not close config file due to error: %v", err)
		}
	}(file)

	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&c)
	if err != nil {
		log.Fatalf("Could not decode config due to error: %v", err)
	}
}
