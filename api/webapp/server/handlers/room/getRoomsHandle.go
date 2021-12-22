package room

import (
	"encoding/json"
	"goReact/webapp/server/handlers/dto"
	"goReact/webapp/server/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetRoomsHandle returns all Rooms
func GetRoomsHandle() httprouter.Handle {
	db := utils.HandlerDbConnection()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT * FROM ROOM")
		if err != nil {
			log.Fatal(err)
		}

		roomsDto := []dto.RoomDto{}

		for rows.Next() {
			room := dto.RoomDto{}
			err := rows.Scan(
				&room.RoomID,
				&room.RoomNumber,
				&room.PetType,
				&room.HotelID)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			roomsDto = append(roomsDto, room)
		}

		json.NewEncoder(w).Encode(roomsDto)
	}
}
