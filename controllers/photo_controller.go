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

type PhotoController interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

func (pc *photoControllerImpl) List(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	photos, err := pc.service.List(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.SuccessWithData("Success retrieving data", photos))
}

func (pc *photoControllerImpl) Update(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	var request requests.PhotoUpdateRequest
	if err := helpers.Binding(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	if isValid, err := govalidator.ValidateStruct(request); !isValid {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	request.UserID = int(id.(float64))
	request.ID = c.Param("photoId")
	photo, err := pc.service.Update(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.SuccessWithData("Photo Updated Successfully", configs.ResponseObj{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt.Format(configs.TimeFormat),
	}))
}

func (pc *photoControllerImpl) Delete(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	if err := pc.service.Delete(id, c.Param("photoId")); err != nil {
		c.JSON(http.StatusInternalServerError, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.Success("Your photo has been successfully deleted"))
}
