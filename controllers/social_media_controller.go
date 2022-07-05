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

type SocialMediaController interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type socialMediaControllerImpl struct {
	service  services.SocialMediaService
	response helpers.Response
}

func NewSocialMediaController(db *gorm.DB, response helpers.Response) SocialMediaController {
	return &socialMediaControllerImpl{
		service:  services.NewSocialMediaService(db),
		response: helpers.NewResponse(),
	}
}

func (pc *socialMediaControllerImpl) Create(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	var socialMedia models.SocialMedia
	if err := helpers.Binding(c, &socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	socialMedia.UserID = uint(id.(float64))
	if err := pc.service.Create(socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, pc.response.SuccessWithData("SocialMedia Created Successfully", configs.ResponseObj{
		"id":              socialMedia.ID,
		"name":            socialMedia.Name,
		"socialMedia_url": socialMedia.SocialMediaUrl,
		"user_id":         socialMedia.UserID,
		"created_at":      socialMedia.CreatedAt.Format(configs.TimeFormat),
	}))
}

func (pc *socialMediaControllerImpl) List(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	socialMedias, err := pc.service.List(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.SuccessWithData("Success retrieving data", socialMedias))
}

func (pc *socialMediaControllerImpl) Update(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	var request requests.SocialMediaUpdateRequest
	if err := helpers.Binding(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	if isValid, err := govalidator.ValidateStruct(request); !isValid {
		c.JSON(http.StatusBadRequest, pc.response.Error(err.Error()))
		return
	}

	request.UserID = int(id.(float64))
	request.ID = c.Param("socialMediaId")
	socialMedia, err := pc.service.Update(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.SuccessWithData("SocialMedia Updated Successfully", configs.ResponseObj{
		"id":              socialMedia.ID,
		"name":            socialMedia.Name,
		"socialMedia_url": socialMedia.SocialMediaUrl,
		"user_id":         socialMedia.UserID,
		"updated_at":      socialMedia.UpdatedAt.Format(configs.TimeFormat),
	}))
}

func (pc *socialMediaControllerImpl) Delete(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, pc.response.Error("Unauthorized"))
		return
	}

	if err := pc.service.Delete(id, c.Param("socialMediaId")); err != nil {
		c.JSON(http.StatusInternalServerError, pc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pc.response.Success("Your social media has been successfully deleted"))
}
