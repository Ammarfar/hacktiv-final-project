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

type CommentController interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CommentControllerImpl struct {
	service  services.CommentService
	response helpers.Response
}

func NewCommentController(db *gorm.DB, response helpers.Response) CommentController {
	return &CommentControllerImpl{
		service:  services.NewCommentService(db),
		response: helpers.NewResponse(),
	}
}

func (cc *CommentControllerImpl) Create(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, cc.response.Error("Unauthorized"))
		return
	}

	var comment models.Comment
	if err := helpers.Binding(c, &comment); err != nil {
		c.JSON(http.StatusBadRequest, cc.response.Error(err.Error()))
		return
	}

	comment.UserID = uint(id.(float64))
	if err := cc.service.Create(comment); err != nil {
		c.JSON(http.StatusBadRequest, cc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, cc.response.SuccessWithData("Comment Created Successfully", configs.ResponseObj{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt.Format(configs.TimeFormat),
	}))
}

func (cc *CommentControllerImpl) List(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, cc.response.Error("Unauthorized"))
		return
	}

	comments, err := cc.service.List(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, cc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, cc.response.SuccessWithData("Success retrieving data", comments))
}

func (cc *CommentControllerImpl) Update(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, cc.response.Error("Unauthorized"))
		return
	}

	var request requests.CommentUpdateRequest
	if err := helpers.Binding(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, cc.response.Error(err.Error()))
		return
	}

	if isValid, err := govalidator.ValidateStruct(request); !isValid {
		c.JSON(http.StatusBadRequest, cc.response.Error(err.Error()))
		return
	}

	request.UserID = int(id.(float64))
	request.ID = c.Param("commentId")
	comment, err := cc.service.Update(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, cc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, cc.response.SuccessWithData("Comment Updated Successfully", configs.ResponseObj{
		"id":         comment.ID,
		"message":    comment.Message,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt.Format(configs.TimeFormat),
	}))
}

func (cc *CommentControllerImpl) Delete(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, cc.response.Error("Unauthorized"))
		return
	}

	if err := cc.service.Delete(id, c.Param("commentId")); err != nil {
		c.JSON(http.StatusInternalServerError, cc.response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, cc.response.Success("Your comment has been successfully deleted"))
}
