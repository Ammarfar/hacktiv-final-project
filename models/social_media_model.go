package models

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaUrl string
	UserID         uint
}
