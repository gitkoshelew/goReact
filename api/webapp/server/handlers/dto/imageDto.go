package dto

// ImageDto ...
type ImageDto struct {
	ImageID int    `json:"imageId"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	OwnerID int    `json:"ownerId"`
}
