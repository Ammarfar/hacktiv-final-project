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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	var checkDuplicate *User
	if tx.Where("username = ?", u.Username).Or("email = ?", u.Email).First(&checkDuplicate).Error == nil {
		err = errors.New("Username or Email has been registered")
		return
	}

	if u.Age < 9 {
		err = errors.New("Age must be greater than 8")
		return
	}

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password, err = helpers.HashPass(u.Password)
	if err != nil {
		return
	}

	err = nil
	return
}
