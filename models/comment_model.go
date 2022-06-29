package models

type Comment struct {
	Common
	UserID  uint
	PhotoID uint
	Message string
}
