package controllers

import (
	"finalproject/helpers"
	"finalproject/models"
	"finalproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c *gin.Context)
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
	helpers.Binding(c, &user)

	if err := uc.service.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, uc.response.SuccessWithData("User Created Successfully", user.Username))
}
