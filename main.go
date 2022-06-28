package main

import (
	"finalproject/configs"
	"finalproject/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	host = helpers.GetEnv("APP_HOST")
	port = helpers.GetEnv("APP_PORT")
)

func main() {

	configs.ConnectDB()
	db := configs.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(host + ":" + port)
}
