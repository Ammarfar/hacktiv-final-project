package controllers

import (
	"finalproject/helpers"
	"finalproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c *gin.Context)
}

type userControllerImpl struct {
	user     *models.User
	db       *gorm.DB
	response helpers.Response
}

func NewUserController(db *gorm.DB, response helpers.Response) UserController {
	return &userControllerImpl{
		user:     &models.User{},
		db:       db,
		response: response,
	}
}

func (this *userControllerImpl) Register(c *gin.Context) {
	helpers.Binding(c, this.user)

	if err := this.db.Debug().Create(&this.user).Error; err != nil {
		c.JSON(http.StatusBadRequest, this.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, this.response.SuccessWithData("User Created Successfully", &this.user.Username))
	return
}
