package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Common
	Title    string    `json:"title" valid:"required"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url" valid:"required"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `json:"comments"`
	User     *User
}

func (u *Photo) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	return nil
}
