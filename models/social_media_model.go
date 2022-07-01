package models

type SocialMedia struct {
	Common
	Name           string `json:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required"`
	UserID         uint   `json:"user_id"`
}
