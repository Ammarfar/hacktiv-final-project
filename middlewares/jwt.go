package middlewares

import (
	"finalproject/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(c *gin.Context) {
	response := helpers.NewResponse()
	errorMessage := "Unauthorized"
	headerToken := c.GetHeader("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		c.JSON(http.StatusUnauthorized, response.Error(errorMessage))
		c.Abort()
		return
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, err := jwt.ParseWithClaims(stringToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(errorMessage)
		}
		return []byte(helpers.GetEnv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, response.Error(errorMessage))
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Set("id", claims["id"])
	c.Writer.Header().Set("Authorization", "Bearer "+stringToken)
	c.Next()
}
