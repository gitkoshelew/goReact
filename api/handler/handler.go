package handler

import (
	"net/http"
)

// New ...
func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	return withLogger(mux)
}
