package models

type Photo struct {
	Common
	Title    string    `json:"title" valid:"required"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url" valid:"required"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `json:"comments"`
}
