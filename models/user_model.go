package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string
	Password     string
	Age          int
	Photos       []Photo
	SocialMedias []SocialMedia
}
