package oauth_test

import (
	"goReact/domain/reqAndResp/oauth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *oauth.GitSSOUser
		isValid bool
	}{
		{name: "valid",
			u: func() *oauth.GitSSOUser {
				return oauth.TestGitUser()
			},
			isValid: true,
		},
		{name: "empty email",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{name: "invalid email",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.Email = "324234"
				return u
			},
			isValid: false,
		},
		{name: "SQL email",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.Email = "Sel--%^3 /** ecT"
				return u
			},
			isValid: false,
		},
		{name: "empty userid",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.UserID = 0
				return u
			},
			isValid: false,
		},
		{name: "empty name",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{name: "SQL name",
			u: func() *oauth.GitSSOUser {
				u := oauth.TestGitUser()
				u.Name = "AlT*&^er"
				return u
			},
			isValid: false,
		},
		{name: "SQL photo",
		u: func() *oauth.GitSSOUser {
			u := oauth.TestGitUser()
			u.Photo = "AlT*&^er"
			return u
		},
		isValid: false,
	},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
