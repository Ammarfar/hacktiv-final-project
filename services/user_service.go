package services

import (
	"errors"
	"finalproject/helpers"
	"finalproject/models"
	"finalproject/repositories"
	"finalproject/requests"
	"fmt"

	"gorm.io/gorm"
)

type UserService interface {
	Register(user models.User) error
	Login(request requests.LoginRequest) (string, error)
	Update(request requests.UserUpdateRequest) (*models.User, error)
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

func (us *userServiceImpl) Login(request requests.LoginRequest) (string, error) {
	user, err := us.repository.GetUserByEmail(request)
	if err != nil {
		return "", errors.New("wrong username or password")
	}

	if success := helpers.ComparePass(user.Password, request.Password); !success {
		return "", errors.New("wrong username or password")
	}

	token, errToken := user.GenerateToken()
	if errToken != nil {
		return "", errors.New("failed generating token")
	}

	return token, nil
}

func (us *userServiceImpl) Update(request requests.UserUpdateRequest) (*models.User, error) {
	if isExist, err := us.repository.IsEmailExist(request); isExist {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("email already exist")
	}

	user, err := us.repository.Update(request)
	if err != nil {
		return nil, err
	}

	return user, nil
}
