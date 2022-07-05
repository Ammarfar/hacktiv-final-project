package responses

type CommentListResponse struct {
	ID        int                      `json:"id"`
	Message   string                   `json:"message"`
	PhotoID   uint                     `json:"photo_id"`
	UserID    uint                     `json:"user_id"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt string                   `json:"updated_at"`
	User      CommentListUserResponse  `json:"User"`
	Photo     CommentListPhotoResponse `json:"Photo"`
}

type CommentListUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentListPhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}
