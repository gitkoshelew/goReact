package service

import (
	"encoding/json"
	"goReact/domain/entity"
	"io/ioutil"
	"log"
)

var hotel entity.Hotel

func loadHotelData() {
	data, err := ioutil.ReadFile("mock/hotel.json")
	if err != nil {
		log.Fatalf("Could not read hotel file due to error: %v", err)
	}

	err = json.Unmarshal(data, &hotel)
	if err != nil {
		log.Fatalf("Could not unmarshal hotel json due to error: %v", err)
	}
}

func LoadTestData() {
	loadHotelData()
}
