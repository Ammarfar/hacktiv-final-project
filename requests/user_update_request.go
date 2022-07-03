package requests

type UserUpdateRequest struct {
	ID       int
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required"`
}
