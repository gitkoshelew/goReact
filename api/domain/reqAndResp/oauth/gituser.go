package oauth

import (
	"goReact/domain/model"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

// Validate ...
func (u *GitSSOUser) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.By(model.IsSQL)),
		validation.Field(&u.Name, validation.Required, validation.By(model.IsSQL)),
		validation.Field((&u.UserID), validation.Required),
		validation.Field(&u.Photo, validation.By(model.IsSQL)),
	)
}
