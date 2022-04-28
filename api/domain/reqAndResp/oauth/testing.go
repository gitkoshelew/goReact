package oauth

//TestGitUser
func TestGitUser() *GitSSOUser {
	return &GitSSOUser{
		UserID: 1233213,
		Email:  "mail@gmail.com",
		Name:   "John",
		Photo:  "git.photo.com",
	}
}

//TestLinkedInSSOUser
func TestLinkedInSSOUser() *LinkedInSSOUser {
	var m = map[string]interface{}{
		"urn:": "photo",
	}
	return &LinkedInSSOUser{
		UserID:  "dgferg234",
		Email:   "mail@gmail.com",
		Name:    "John",
		Surname: "Smith",
		Photo:   m,
	}
}
