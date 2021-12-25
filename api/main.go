package main

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp"
	"goReact/webapp/server"
	"goReact/webapp/server/utils"
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
	webapp.ConnectDb(config)
	initAccountDb()

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

// initAccountDb hashes accounts password from  DB, that creates by ini.sql
func initAccountDb() {
	db := utils.HandlerDbConnection()
	for i := 0; i < 6; i++ {

		encryptedPassword, err := store.EncryptPassword(fmt.Sprintf("password%d", i+1))
		if err != nil {
			log.Fatal(err)
		}

		result, err := db.Exec("UPDATE ACCOUNT set password = $1 WHERE id = $2",
			encryptedPassword, i+1)
		if err != nil {
			panic(err)
		}
		log.Println(result.RowsAffected())
	}

}
