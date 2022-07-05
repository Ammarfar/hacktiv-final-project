package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Common
	Name           string `json:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required"`
	UserID         uint   `json:"user_id"`
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(sm); err != nil {
		return err
	}

	return nil
}
