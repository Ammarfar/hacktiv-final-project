package main

import (
	"finalproject/configs"
	"finalproject/helpers"
	"finalproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// DB
	configs.ConnectDB()
	db := configs.GetDB()
	sqlDB, err := db.DB()
	helpers.PanicIfError(err)
	defer sqlDB.Close()

	// declaration
	host := helpers.GetEnv("APP_HOST")
	port := helpers.GetEnv("APP_PORT")
	response := helpers.NewResponse()
	server := gin.Default()

	// routes
	routes.Index(server, db, response)

	server.Run(host + ":" + port)
}
