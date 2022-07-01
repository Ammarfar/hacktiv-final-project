package middlewares

import (
	"finalproject/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsMatchUser(c *gin.Context) {
	response := helpers.NewResponse()
	errorMessage := "Unauthorized"

	id, exist := c.Get("id")
	if !exist || c.Param("userId") != strconv.Itoa(int(id.(float64))) {
		c.JSON(http.StatusUnauthorized, response.Error(errorMessage))
		c.Abort()
		return
	}

	c.Next()
}
