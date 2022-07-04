package controllers

import (
	"finalproject/configs"
	"finalproject/helpers"
	"finalproject/models"
	"finalproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController interface {
	Create(c *gin.Context)
}

type photoControllerImpl struct {
	service  services.PhotoService
	response helpers.Response
}

func NewPhotoController(db *gorm.DB, response helpers.Response) PhotoController {
	return &photoControllerImpl{
		service:  services.NewPhotoService(db),
		response: helpers.NewResponse(),
	}
}

func (pc *photoControllerImpl) Create(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	var photo models.Photo
	if err := helpers.Binding(c, &photo); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	photo.UserID = uint(id.(float64))
	if err := pc.service.Create(photo); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, pc.response.SuccessWithData("Photo Created Successfully", configs.ResponseObj{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt.Format(configs.TimeFormat),
	}))
}
