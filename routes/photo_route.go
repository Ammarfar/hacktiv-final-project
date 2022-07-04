package routes

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRoute(r *gin.Engine, controller controllers.PhotoController) {
	photos := r.Group("/photos", middlewares.VerifyToken)
	{
		photos.POST("/", controller.Create)
		photos.GET("/", controller.List)
		photos.PUT("/:photoId", controller.Update)
		photos.DELETE("/:photoId", controller.Delete)
	}
}
