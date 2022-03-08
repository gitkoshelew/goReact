package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIError ...
type APIError struct {
	Message          string `json:"message,omitempty"`
	ErrorCode        string `json:"error_code,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
}

// APIResponse ...
type APIResponse struct {
	IsOk     bool
	response *http.Response
	Error    APIError
}

// ReadBody ...
func (r *APIResponse) ReadBody() ([]byte, error) {
	defer r.response.Body.Close()
	return ioutil.ReadAll(r.response.Body)
}

// ToString ...
func (r *APIError) ToString() string {
	return fmt.Sprintf("Err Code: %s, Err: %s, Developer Err: %s", r.ErrorCode, r.Message, r.DeveloperMessage)
}

// IsResponseOk ...
func (r *APIResponse) IsResponseOk() {
	r.IsOk = true

	if r.response.StatusCode < http.StatusOK || r.response.StatusCode >= http.StatusBadRequest {
		r.IsOk = false
		defer r.response.Body.Close()

		var apiErr APIError
		if err := json.NewDecoder(r.response.Body).Decode(&apiErr); err == nil {
			r.Error = apiErr
		}
	}
}
