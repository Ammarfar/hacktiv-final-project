package services

import (
	"finalproject/models"
	"finalproject/repositories"
	"fmt"

	"gorm.io/gorm"
)

type UserService interface {
	Register(user models.User) error
}

type userServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		repository: repositories.NewUserRepository(db),
	}
}

func (us *userServiceImpl) Register(request models.User) error {
	if us.repository.IsUserExist(request) {
		return fmt.Errorf("username %s or email %s has been registered", request.Username, request.Email)
	}

	return us.repository.Register(request)
}
