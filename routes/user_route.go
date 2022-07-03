package routes

import (
	"finalproject/controllers"
	"finalproject/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, controller controllers.UserController) {
	users := r.Group("/users")
	{
		users.POST("/register", controller.Register)
		users.POST("/login", controller.Login)
		users.PUT("/:userId", middlewares.VerifyToken, middlewares.IsMatchUser, controller.UpdateUser)
		users.DELETE("/", middlewares.VerifyToken, controller.Delete)
	}
}
