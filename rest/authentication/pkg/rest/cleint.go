package rest

import (
	"errors"
	"fmt"
	"gateway/pkg/logging"
	"net/http"
	"net/url"
	"path"
)

// BaseClient ...
type BaseClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Logger     *logging.Logger
}

var (
	// ErrNoHTTPClient ...
	ErrNoHTTPClient = errors.New("no http client")
)

// SendRequest ...
func (c *BaseClient) SendRequest(req *http.Request) (*APIResponse, error) {
	if c.HTTPClient == nil {
		return nil, ErrNoHTTPClient
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request. error: %w", err)
	}

	apiResponse := APIResponse{
		response: response,
	}

	apiResponse.IsResponseOk()

	return &apiResponse, nil
}

// BuildURL ...
func (c *BaseClient) BuildURL(endpoint string, filters []FilterOptions) (string, error) {

	parsedURL, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL. error: %w", err)
	}

	parsedURL.Path = path.Join(parsedURL.Path, endpoint)

	if len(filters) > 0 {
		q := parsedURL.Query()
		for _, fo := range filters {
			q.Set(fo.Field, fo.ToString())
		}
		parsedURL.RawQuery = q.Encode()
	}

	return parsedURL.String(), nil
}

// Close ...
func (c *BaseClient) Close() error {
	c.HTTPClient = nil
	return nil
}
