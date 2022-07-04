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
	}
}
