package main

import (
	"database/sql"
	"fmt"
	"goReact/domain/model"
	"goReact/webapp"
	"goReact/webapp/server"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Need this part? If we have .env in docker-compose.yaml
	err := godotenv.Load(filepath.Join("/", ".env"))
	if err != nil {
		log.Printf("Error loading .env file. %s", err.Error())
	}

	config := &webapp.Config{}
	config.NewConfig()
	initAccountDb(config)
	// createRooms(config, 100)

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

// initAccountDb hashes accounts password from  DB, that creates by ini.sql
func initAccountDb(c *webapp.Config) {
	dataSourceName := c.PgDataSource()

	log.Printf("Connecting to database via %#v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for i := 0; i < 6; i++ {

		encryptedPassword, err := model.EncryptPassword(fmt.Sprintf("password"))
		if err != nil {
			log.Print(err)
		}

		_, err = db.Exec("UPDATE USERS set password = $1 WHERE id = $2",
			encryptedPassword, i+1)
		if err != nil {
			log.Print(err)
		}
	}
	log.Println("passwords encrypted")

}

// func createRooms(c *webapp.Config, roomCount int) {

// 	var petType model.PetType
// 	var hotelID int

// 	for i := 0; i < roomCount; i++ {

// 		dataSourceName := c.PgDataSource()
// 		db, err := sql.Open("postgres", dataSourceName)
// 		if err != nil {
// 			panic(err)
// 		}

// 		err = db.Ping()
// 		if err != nil {
// 			panic(err)
// 		}

// 		if i%2 == 0 {
// 			petType = model.PetTypeDog
// 		} else {
// 			petType = model.PetTypeCat
// 		}

// 		rand.Seed(time.Now().UnixNano())
// 		n := rand.Int63()
// 		if n%2 == 0 {
// 			hotelID = 1
// 		} else {
// 			hotelID = 2
// 		}

// 		if err := db.QueryRow(
// 			"INSERT INTO room (pet_type, number, hotel_id) VALUES ($1, $2, $3)",
// 			string(petType), i, hotelID); err != nil {

// 			log.Printf("err %v", err)
// 		}

// 		db.Close()
// 	}
// 	log.Printf("%d rooms added to DB", roomCount)
// }
