package routes

import (
	"finalproject/controllers"
	"finalproject/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(r *gin.Engine, db *gorm.DB, response helpers.Response) {
	UserRoute(r, controllers.NewUserController(db, response))
}
