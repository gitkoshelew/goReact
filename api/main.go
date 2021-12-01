package main

import (
	"goReact/webapp"
	"goReact/webapp/server"
	"log"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())

	}
	config := &webapp.Config{}
	config.ReadFromFile("config.yaml")

	s := server.New(config)
	/*if err := s.Run(viper.GetString("8081"), handlers.InitRoutes()); err != nil{

	}*/

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
