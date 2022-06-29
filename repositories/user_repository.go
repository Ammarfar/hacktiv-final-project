package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (ur *userRepositoryImpl) Register(user models.User) (*models.User, error) {

	if err := ur.db.Debug().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
