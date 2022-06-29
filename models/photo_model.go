package models

type Photo struct {
	Common
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
	Comments []Comment
}
