package repositories

import (
	"finalproject/models"
	"finalproject/requests"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) error
	IsUserExist(user models.User) bool
	GetUserByEmail(request requests.LoginRequest) (*models.User, error)
	Update(request requests.UserUpdateRequest) (*models.User, error)
	IsEmailExist(request requests.UserUpdateRequest) (bool, error)
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

func (ur *userRepositoryImpl) GetUserByEmail(request requests.LoginRequest) (*models.User, error) {
	var user models.User
	if err := ur.db.Select("id", "password").Where("email = ?", request.Email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepositoryImpl) Update(request requests.UserUpdateRequest) (*models.User, error) {
	var user models.User

	if err := ur.db.Where("id = ?", request.ID).Take(&user).Error; err != nil {
		return nil, err
	}

	user.Username = request.Username
	user.Email = request.Email

	if err := ur.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepositoryImpl) IsEmailExist(request requests.UserUpdateRequest) (bool, error) {
	var user models.User

	if err := ur.db.Where("email = ?", request.Email).Where("id != ?", request.ID).Take(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
