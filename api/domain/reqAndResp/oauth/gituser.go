package oauth

import (
	"goReact/domain/model"
	"strconv"
	"time"
)

//GitSSOUser response
type GitSSOUser struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Photo  string `json:"avatar_url"`
}

//User model from gituser
func UserFromGit(gituser *GitSSOUser) (*model.User, error) {
	verified := new(bool)
	*verified = false
	dateOfBirth := time.Time{}

	return &model.User{
		Email:           gituser.Email,
		Name:            gituser.Name,
		Surname:         gituser.Name,
		Sex:             model.SexUnknown,
		DateOfBirth:     &dateOfBirth,
		Photo:           gituser.Photo,
		SocialNetwork:   model.GitHub,
		SocialNetworkID: strconv.Itoa(gituser.UserID),
	}, nil

}
