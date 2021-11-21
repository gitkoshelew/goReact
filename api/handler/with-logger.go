package handler

import (
	"log"
	"net/http"
	"os"
)

func withLogger(next http.Handler) http.Handler {
	logger := log.New(os.Stdout, "http: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			logger.Println(req.Method, req.URL.Path)
		}()
		next.ServeHTTP(res, req)
	})
}
