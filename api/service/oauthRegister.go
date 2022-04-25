package service

import (
	"goReact/domain/model"
	"goReact/domain/store"
)

//Checks if there is a user, if not, then registers it
func OauthRegisterUser(s *store.Store, newUser *model.User) (*int, error) {
	err := s.Open()
	if err != nil {
		return nil, err
	}

	user, err := s.User().FindBySocialNetworkId(newUser.SocialNetworkID)
	if err != nil {
		_, err = s.User().CreateFromSocial(newUser)
		if err != nil {
			return nil, err
		}
		return &newUser.UserID, nil
	}

	return &user.UserID, nil
}
