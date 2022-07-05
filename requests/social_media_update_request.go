package requests

type SocialMediaUpdateRequest struct {
	ID             string
	UserID         int
	Name           string `json:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required"`
}
