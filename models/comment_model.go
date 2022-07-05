package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Common
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" valid:"required"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return nil
}
