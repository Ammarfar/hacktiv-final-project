package controllers

import (
	"finalproject/configs"
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
	UpdateUser(c *gin.Context)
	Delete(c *gin.Context)
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

	c.JSON(http.StatusCreated, uc.response.SuccessWithData("User Created Successfully", configs.ResponseObj{
		"age":     user.Age,
		"email":   user.Email,
		"id":      user.ID,
		"usernam": user.Username,
	}))
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

	c.JSON(http.StatusOK, uc.response.SuccessWithData("Login Success", configs.ResponseObj{"token": token}))
}

func (uc *userControllerImpl) UpdateUser(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, uc.response.Error("Unauthorized"))
		return
	}

	var request requests.UserUpdateRequest
	if err := helpers.Binding(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	if isValid, err := govalidator.ValidateStruct(request); !isValid {
		c.JSON(http.StatusBadRequest, uc.response.Error(err.Error()))
		return
	}

	request.ID = int(id.(float64))
	user, err := uc.service.Update(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, uc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, uc.response.SuccessWithData("User Updated Successfully", configs.ResponseObj{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt.Format(configs.TimeFormat),
	}))
}

func (uc *userControllerImpl) Delete(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, uc.response.Error("Unauthorized"))
		return
	}

	if err := uc.service.Delete(int(id.(float64))); err != nil {
		c.JSON(http.StatusInternalServerError, uc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, uc.response.Success("Your account has been successfully deleted"))
}
