package routes

import (
	"finalproject/controllers"
	"finalproject/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(r *gin.Engine, db *gorm.DB) {
	response := helpers.NewResponse()

	UserRoute(r, controllers.NewUserController(db, response))
	PhotoRoute(r, controllers.NewPhotoController(db, response))
	CommentRoute(r, controllers.NewCommentController(db, response))
	SocialMediaRoute(r, controllers.NewSocialMediaController(db, response))
}
