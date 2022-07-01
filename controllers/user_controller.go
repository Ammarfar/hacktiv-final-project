package controllers

import (
	"finalproject/helpers"
	"finalproject/models"
	"finalproject/requests"
	"finalproject/services"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userControllerImpl struct {
	service  services.UserService
	response helpers.Response
}

func NewUserController(db *gorm.DB, response helpers.Response) UserController {
	return &userControllerImpl{
		service:  services.NewUserService(db),
		response: response,
	}
}

func (uc *userControllerImpl) Register(c *gin.Context) {
	var user models.User
	if err := helpers.Binding(c, &user); err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	if err := uc.service.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, uc.response.SuccessWithData("User Created Successfully", user.Username))
}

func (uc *userControllerImpl) Login(c *gin.Context) {
	var request requests.LoginRequest
	if err := helpers.Binding(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	if isValid, err := govalidator.ValidateStruct(request); !isValid {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	token, err := uc.service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, uc.response.SuccessWithData("Login Success", helpers.ResponseObj{"token": token}))
}
