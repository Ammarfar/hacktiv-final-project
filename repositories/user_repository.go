package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) error
	IsUserExist(user models.User) bool
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (ur *userRepositoryImpl) Register(user models.User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepositoryImpl) IsUserExist(user models.User) bool {
	return ur.db.Where("username = ?", user.Username).Or("email = ?", user.Email).First(&user).Error == nil
}
