package services

import (
	"finalproject/models"
	"finalproject/repositories"

	"gorm.io/gorm"
)

type UserService interface {
	Register(user models.User) (*models.User, error)
}

type userServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		repository: repositories.NewUserRepository(db),
	}
}

func (us *userServiceImpl) Register(request models.User) (*models.User, error) {

	user, err := us.repository.Register(request)

	if err != nil {
		return nil, err
	}

	return user, nil
}
