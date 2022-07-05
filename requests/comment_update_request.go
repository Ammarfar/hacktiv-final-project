package requests

type CommentUpdateRequest struct {
	ID      string
	UserID  int
	Message string `json:"message" valid:"required"`
}
