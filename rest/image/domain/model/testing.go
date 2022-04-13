package model

// TestImage ...
func TestImage() *Image {
	return &Image{
		Type:    UserImage,
		OwnerID: 1,
		Format:  "original",
	}
}

// TestImageDTO ...
func TestImageDTO() *ImageDTO {
	return &ImageDTO{
		Type:    string(TestingImage),
		OwnerID: 1,
		Format:  "original",
	}
}
