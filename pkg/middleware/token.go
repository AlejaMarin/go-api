package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	requiredToken := os.Getenv("TOKEN")
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "token not found"})
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}
		c.Next()
	}
}
