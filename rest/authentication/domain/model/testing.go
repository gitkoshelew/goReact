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
