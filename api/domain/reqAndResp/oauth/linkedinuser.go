package oauth

import (
	"fmt"
	"goReact/domain/model"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//LinkedInSSOUser response
type LinkedInSSOUser struct {
	UserID  string                 `json:"id"`
	Email   string                 `json:"email"`
	Name    string                 `json:"localizedFirstName"`
	Surname string                 `json:"localizedLastName"`
	Photo   map[string]interface{} `json:"profilePicture"`
}

//User model from linkedinuser
func UserFromLinked(linkedInUser *LinkedInSSOUser) (*model.User, error) {
	dateOfBirth := time.Time{}

	return &model.User{
		Name:            linkedInUser.Name,
		Surname:         linkedInUser.Surname,
		Sex:             model.SexUnknown,
		DateOfBirth:     &dateOfBirth,
		Photo:           fmt.Sprintf("%v", linkedInUser.Photo),
		SocialNetwork:   model.LinkedIN,
		SocialNetworkID: linkedInUser.UserID,
	}, nil

}

// Validate ...
func (u *LinkedInSSOUser) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, is.Email, validation.By(model.IsSQL)),
		validation.Field(&u.Name, validation.Required, validation.By(model.IsSQL)),
		validation.Field(&u.Surname, validation.Required, validation.By(model.IsSQL)),
		validation.Field(&u.UserID, validation.Required, validation.By(model.IsSQL)),
	)
}
