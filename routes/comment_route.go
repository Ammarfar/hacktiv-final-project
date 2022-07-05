package routes

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoute(r *gin.Engine, controller controllers.CommentController) {
	comments := r.Group("/comments", middlewares.VerifyToken)
	{
		comments.POST("/", controller.Create)
		comments.GET("/", controller.List)
		comments.PUT("/:commentId", controller.Update)
		comments.DELETE("/:commentId", controller.Delete)
	}
}
