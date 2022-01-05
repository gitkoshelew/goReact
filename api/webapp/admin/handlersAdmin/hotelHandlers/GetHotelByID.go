package hotelhandlers

import (
	"fmt"
	"goReact/domain/model"
	"goReact/webapp/server/utils"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GetHotelByID ...
func GetHotelByID() httprouter.Handle {
	db := utils.HandlerDbConnection()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		hotels := []model.Hotel{}

		id, _ := strconv.Atoi(ps.ByName("id"))
		rows, err := db.Query("select * from hotel where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			h := model.Hotel{}
			err := rows.Scan(&h.HotelID, &h.Name, &h.Address)
			if err != nil {
				fmt.Println(err)
				continue
			}
			hotels = append(hotels, h)
		}

		if len(hotels) == 0 {
			http.Error(w, "No hotel with such id!", 400)
			return
		}

		files := []string{
			"/api/webapp/admin/tamplates/allHotels.html",
			"/api/webapp/admin/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, hotels)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
