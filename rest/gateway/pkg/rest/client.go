package rest

import (
	"context"
	"errors"
	"gateway/pkg/logging"
	"gateway/pkg/response"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
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
	// ErrNoHeaders ...
	ErrNoHeaders = errors.New("no headers")
	// ErrNoCookies ...
	ErrNoCookies = errors.New("no cookies. Sad... :( ")
)

// BuildURL ...
func (c *BaseClient) BuildURL(endpoint string, filters []FilterOptions) (string, error) {
	c.Logger.Debug("build url with resource")
	parsedURL, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		c.Logger.Errorf("failed to parse base URL. error: %w", err)
		return "", err
	}

	parsedURL.Path = path.Join(parsedURL.Path, endpoint)

	if len(filters) > 0 {
		q := parsedURL.Query()
		for _, fo := range filters {
			q.Set(fo.Field, fo.ToString())
		}
		parsedURL.RawQuery = q.Encode()
	}
	c.Logger.Debugf("URL: %v", parsedURL)
	return parsedURL.String(), nil
}

// CreateRequest ...
// default headers: "Accept", "application/json"
// 					"Content-Type", "application/json"
func (c *BaseClient) CreateRequest(method, url string, body io.Reader) (*http.Request, error) {
	c.Logger.Debug("create new request")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.Logger.Errorf("failed to parse base URL. error: %w", err)
		return nil, err
	}

	if c.HTTPClient == nil {
		return nil, ErrNoHTTPClient
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// SendRequest ...
func (c *BaseClient) SendRequest(ctx context.Context, req *http.Request) (*APIResponse, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "image/jpeg")

	ctxTimeOut, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	req = req.WithContext(ctxTimeOut)
	c.Logger.Debug("sending request...")
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Errorf("failed to send request. error: %w", err)
		return nil, err
	}

	apiResponse := APIResponse{
		response: response,
	}

	apiResponse.IsResponseOk()
	return &apiResponse, nil
}

// ReadResponse returns http.response with body, cokkies and headers.
// headers keys need to be send with function call
func (c *BaseClient) ReadResponse(resp *APIResponse, headersKeys []string) (*response.Service, error) {
	c.Logger.Debug("read body")
	tags, err := resp.ReadBody()
	if err != nil {
		c.Logger.Infof("failed to read body. err: %w", err)
	}

	c.Logger.Debug("read headers")
	headers, err := resp.ReadHeaders(headersKeys)
	if err != nil {
		c.Logger.Infof("no headers in response")
	}

	c.Logger.Debug("read cookies")
	cookies, err := resp.ReadCookies()
	if err != nil {
		c.Logger.Infof("no cookies in response")
	}

	return &response.Service{
		Body:    tags,
		Headers: headers,
		Cookies: cookies,
	}, nil
}

// Close ...
func (c *BaseClient) Close() error {
	c.HTTPClient = nil
	return nil
}
