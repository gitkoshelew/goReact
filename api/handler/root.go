package handler

import "net/http"

func handleRoot(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Header().Set("Allow", http.MethodGet)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, _ = res.Write([]byte("{\"app\": \"Go React project\"}"))
}
