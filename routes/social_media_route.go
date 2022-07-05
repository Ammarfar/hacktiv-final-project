package routes

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func SocialMediaRoute(r *gin.Engine, controller controllers.SocialMediaController) {
	socialMedias := r.Group("/socialmedias", middlewares.VerifyToken)
	{
		socialMedias.POST("/", controller.Create)
		socialMedias.GET("/", controller.List)
		socialMedias.PUT("/:socialMediaId", controller.Update)
		socialMedias.DELETE("/:socialMediaId", controller.Delete)
	}
}
