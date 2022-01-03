package handlers

import (
	"net/http"
)

// CorsHandle ...
func CorsHandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
