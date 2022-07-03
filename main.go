package main

import (
	"finalproject/configs"
	"finalproject/controllers"
	"finalproject/helpers"
	"finalproject/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	host     = helpers.GetEnv("APP_HOST")
	port     = helpers.GetEnv("APP_PORT")
	response = helpers.NewResponse()
)

func main() {

	configs.ConnectDB()
	db := configs.GetDB()
	sqlDB, err := db.DB()
	helpers.PanicIfError(err)
	defer sqlDB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userController := controllers.NewUserController(db, response)
	users := r.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.PUT("/:userId", middlewares.VerifyToken, middlewares.IsMatchUser, userController.UpdateUser)
	}

	r.Run(host + ":" + port)
}
