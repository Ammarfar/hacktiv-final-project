package responses

type SocialMediaListResponse struct {
	ID             int                         `json:"id"`
	Name           string                      `json:"name"`
	SocialMediaUrl string                      `json:"social_media_url"`
	UserID         int                         `json:"user_id"`
	CreatedAt      string                      `json:"created_at"`
	UpdatedAt      string                      `json:"updated_at"`
	User           SocialMediaListUserResponse `json:"User"`
}

type SocialMediaListUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
