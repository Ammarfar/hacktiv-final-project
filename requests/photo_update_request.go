package requests

type PhotoUpdateRequest struct {
	ID       string
	UserID   int
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required"`
}
