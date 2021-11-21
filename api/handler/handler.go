package handler

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	return withLogger(mux)
}
