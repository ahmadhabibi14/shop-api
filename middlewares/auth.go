package middlewares

import (
	"net/http"
	"shop-api/config"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := config.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, Please login !!")
			c.Abort()
			return
		}
		c.Next()
	}
}
