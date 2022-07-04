package responses

type PhotoListResponse struct {
	ID        int                   `json:"id"`
	Title     string                `json:"title"`
	Caption   string                `json:"caption"`
	PhotoUrl  string                `json:"photo_url"`
	UserID    int                   `json:"user_id"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
	User      PhotoListUserResponse `json:"User"`
}

type PhotoListUserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
