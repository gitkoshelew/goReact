package oauth_test

import (
	"goReact/domain/reqAndResp/oauth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *oauth.LinkedInSSOUser
		isValid bool
	}{
		{name: "valid",
			u: func() *oauth.LinkedInSSOUser {
				return oauth.TestLinkedInSSOUser()
			},
			isValid: true,
		},
		{name: "invalid email",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Email = "324234"
				return u
			},
			isValid: false,
		},
		{name: "SQL email",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Email = "Sel--%^3 /** ecT"
				return u
			},
			isValid: false,
		},
		{name: "empty userid",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.UserID = ""
				return u
			},
			isValid: false,
		},
		{name: "SQL userid",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.UserID = "AlT*&^er"
				return u
			},
			isValid: false,
		},
		{name: "empty name",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{name: "SQL name",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Name = "AlT*&^er"
				return u
			},
			isValid: false,
		},

		{name: "empty surname",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Surname = ""
				return u
			},
			isValid: false,
		},
		{name: "SQL surname",
			u: func() *oauth.LinkedInSSOUser {
				u := oauth.TestLinkedInSSOUser()
				u.Surname = "AlT*&^er"
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
