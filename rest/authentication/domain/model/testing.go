package model

// TestUser ...
func TestUser() *User {
	verified := true
	return &User{
		UserID:   1,
		Email:    "email@example.org",
		Password: "password",
		Role:     "client",
		Verified: &verified,
	}
}

// TestUserDTO ...
func TestUserDTO() *UserDTO {
	verified := true
	return &UserDTO{
		UserID:   1,
		Email:    "email@example.org",
		Password: "password",
		Role:     "client",
		Verified: &verified,
	}
}

// TestLogin ...
func TestLogin() *Login {
	return &Login{
		Email:    "login@example.org",
		Password: "password",
	}
}
