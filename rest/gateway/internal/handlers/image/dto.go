package image

// ImageDTO struct
type ImageDTO struct {
	ImageID int    `json:"imageId"`
	Type    string `json:"type"`
	OwnerID int    `json:"ownerId"`
	Format  string `json:"format"`
}
