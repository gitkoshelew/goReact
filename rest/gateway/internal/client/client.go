package client

import (
	"fmt"
	"gateway/pkg/logging"
	"gateway/pkg/rest"
	"net/http"
	"os"
	"time"
)

var (
	// AuthLoginService ...
	AuthLoginService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/login")
	// AuthLogoutService ...
	AuthLogoutService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/logout")
	// AuthRegistrationService ...
	AuthRegistrationService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/registration")
	// AuthRefreshService ...
	AuthRefreshService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/refresh")

	// CustomerGetAllUsersService ...
	CustomerGetAllUsersService = New(fmt.Sprintf("http://%s:%s", os.Getenv("USER_SERVER_DOCKER_HOST"), os.Getenv("USER_SERVER_PORT")), "/users")
	// CustomerUserService using methods:
	// "POST" to create
	// "PUT" to update
	// "GET" to get one by id. Using querry params (/user/:id)
	// "DELETE" to delete by id. Using querry params (/user/:id)
	CustomerUserService = New(fmt.Sprintf("http://%s:%s", os.Getenv("USER_SERVER_DOCKER_HOST"), os.Getenv("USER_SERVER_PORT")), "/user")
	// CustomerGetAllPetsService ...
	CustomerGetAllPetsService = New(fmt.Sprintf("http://%s:%s", os.Getenv("USER_SERVER_DOCKER_HOST"), os.Getenv("USER_SERVER_PORT")), "/pets")
	// CustomerPetService using methods:
	// "POST" to create
	// "PUT" to update
	// "GET" to get one by id. Using querry params (/user/:id)
	// "DELETE" to delete by id. Using querry params (/user/:id)
	CustomerPetService = New(fmt.Sprintf("http://%s:%s", os.Getenv("USER_SERVER_DOCKER_HOST"), os.Getenv("USER_SERVER_PORT")), "/pet")

	// BookingGetAllUsersService ...
	BookingGetAllUsersService = New(fmt.Sprintf("http://%s:%s", os.Getenv("BOOKING_SERVER_DOCKER_HOST"), os.Getenv("BOOKING_SERVER_PORT")), "/bookings")
	// BookingUserService using methods:
	// "POST" to create
	// "PUT" to update
	// "GET" to get one by id. Using querry params (/booking/:id)
	// "DELETE" to delete by id. Using querry params (/booking/:id)
	BookingUserService = New(fmt.Sprintf("http://%s:%s", os.Getenv("BOOKING_SERVER_DOCKER_HOST"), os.Getenv("BOOKING_SERVER_PORT")), "/booking")
)

// CtxKey ...
type CtxKey int8

const (
	// AccessTokenCtxKey ...
	AccessTokenCtxKey CtxKey = 1
	// RefreshTokenCtxKey ...
	RefreshTokenCtxKey CtxKey = 2
	// CustomerGetQuerryParamsCtxKey ...
	CustomerGetQuerryParamsCtxKey CtxKey = 3
	// CustomerDeleteQuerryParamsCtxKey ...
	CustomerDeleteQuerryParamsCtxKey CtxKey = 4
	// BookingDeleteQuerryParamsCtxKey ...
	BookingDeleteQuerryParamsCtxKey CtxKey = 5
	// BookingGetQuerryParamsCtxKey ...
	BookingGetQuerryParamsCtxKey CtxKey = 6
)

// Client ...
type Client struct {
	Base     rest.BaseClient
	Resource string
}

// New ...
func New(baseURL string, resource string) *Client {
	return &Client{
		Base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logging.GetLogger(),
		},
		Resource: resource,
	}
}
