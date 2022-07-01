package helpers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnv(param string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(param)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Binding(c *gin.Context, obj any) (err error) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		err = c.ShouldBindJSON(obj)
	} else {
		err = c.ShouldBind(obj)
	}

	return err
}
