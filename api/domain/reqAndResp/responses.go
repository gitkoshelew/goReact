package reqandresp

import (
	"fmt"
	"goReact/domain/model"
	"time"
)

// RoomInfo ...
type RoomInfo struct {
	SeatID    int        `json:"seatId"`
	StartDate *time.Time `json:"start"`
	EndDate   *time.Time `json:"end"`
}

// FreeSeatsResponse struct
type FreeSeatsResponse struct {
	Day        int   `json:"day"`
	SeatIDs    []int `json:"seatIds"`
	TotalCount int   `json:"totalCount"`
}

//Top rooms response
type TopRooms struct {
	RoomID        int `json:"roomId"`
	TotalBookings int `json:"total bookings"`
}

//GitSSOUser response
type GitSSOUser struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Photo  string `json:"avatar_url"`
}

//LinkedInSSOUser response
type LinkedInSSOUser struct {
	UserID  string                 `json:"id"`
	Email   string                 `json:"email"`
	Name    string                 `json:"localizedFirstName"`
	Surname string                 `json:"localizedLastName"`
	Photo   map[string]interface{} `json:"profilePicture"`
}

func UserFromGit(gituser *GitSSOUser) (*model.User, error) {
	return &model.User{
		Name:    gituser.Name,
		Surname: gituser.Name,
		Email:   gituser.Email,
		Role:    model.ClientRole,
		Photo:   gituser.Photo,
	}, nil

}

func UserFromLinked(linkedInUser *LinkedInSSOUser) (*model.User, error) {
	photourl := linkedInUser.Photo["profilePicture"]
	return &model.User{
		Name:    linkedInUser.Name,
		Surname: linkedInUser.Surname,
		Email:   linkedInUser.Email,
		Role:    model.ClientRole,
		Photo:   fmt.Sprintf("%v", photourl),
	}, nil

}
