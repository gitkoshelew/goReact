package response

import "net/http"

// Error ...
type Error struct {
	Messsage string `json:"ErrorMsg"`
}

// Info ...
type Info struct {
	Messsage string `json:"infoMsg"`
}

// Service ...
type Service struct {
	Body    []byte
	Headers map[string]string
	Cookies []*http.Cookie
}
