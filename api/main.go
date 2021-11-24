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

	// config := &webapp.Config{}
	// config.ReadFromFile("config.yaml")

	// service.LoadTestData()
	// db, err := webapp.ConnectDb(config)
	// if err != nil {
	// 	panic(fmt.Errorf("could not initialize db connection: %v", err))
	// }
	// defer db.Close()

	// rows, err := db.Query("SELECT COUNT(*) FROM booking")
	// if err != nil {
	// 	panic("could not execute query")
	// }
	// defer rows.Close()
	// var count int
	// rows.Next()
	// err = rows.Scan(&count)
	// if err != nil {
	// 	panic("could not scan count")
	// }
	// fmt.Printf("Server started on: %d,/n Bookings count: %d", config.Server.Port, count)
	//
	// webapp.StartServer(config)
}
