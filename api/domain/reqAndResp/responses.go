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
	verified := new(bool)
	*verified = false
	dateOfBirth := time.Time{}

	return &model.User{
		Email: gituser.Email,
		//	Password:    "1",
		//Role:        model.ClientRole,
		//	Verified:    verified,
		Name:    gituser.Name,
		Surname: gituser.Name,
		//	MiddleName:  "ААаа",
		Sex:         model.SexUnknown,
		DateOfBirth: &dateOfBirth,
		//Address:     "1",
		//	Phone:       "1",
		Photo:         gituser.Photo,
		SocialNetwork: model.GitHub,
	}, nil

}

func UserFromLinked(linkedInUser *LinkedInSSOUser) (*model.User, error) {
	dateOfBirth := time.Time{}

	return &model.User{
		//Email:       "",
		//Password:    "",
		//Role:        model.ClientRole,
		Name:    linkedInUser.Name,
		Surname: linkedInUser.Surname,
		//	MiddleName:    "",
		Sex:         model.SexUnknown,
		DateOfBirth: &dateOfBirth,
		//Address:       "",
		//Phone:         "",
		Photo:         fmt.Sprintf("%v", linkedInUser.Photo),
		SocialNetwork: model.LinkedIN,
	}, nil

}
