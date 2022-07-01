package models

import (
	"errors"
	"finalproject/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Common
	Username     string        `json:"username" valid:"required"`
	Email        string        `json:"email" valid:"required,email"`
	Password     string        `json:"password" valid:"required,minstringlength(6)~Password minimum is 6 character"`
	Age          int           `json:"age" valid:"required"`
	Photos       []Photo       `json:"photos"`
	SocialMedias []SocialMedia `json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error

	if u.Age < 9 {
		return errors.New("age must be greater than 8")
	}

	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password, err = helpers.HashPass(u.Password)
	if err != nil {
		return err
	}

	return nil
}
