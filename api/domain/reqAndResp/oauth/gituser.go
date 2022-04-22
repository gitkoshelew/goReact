package oauth

import (
	"goReact/domain/model"
	"time"
)

//GitSSOUser response
type GitSSOUser struct {
	UserID string `json:"id"`
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
		SocialNetworkID: gituser.UserID,
	}, nil

}
