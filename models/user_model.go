package models

type User struct {
	Common
	Username     string
	Email        string
	Password     string
	Age          int
	Photos       []Photo
	SocialMedias []SocialMedia
}
