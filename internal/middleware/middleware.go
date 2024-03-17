package middleware

import (
	"cart/internal/confi"

	"github.com/gin-gonic/gin"
)

func BasicAuth(appconfig *confi.AppConfig) gin.HandlerFunc {
	// Create the Gin middleware handler for BasicAuth
	authMiddleware := gin.BasicAuth(gin.Accounts{
		"yogi": "1234",
	})

	return func(c *gin.Context) {
		// Execute the BasicAuth middleware
		authMiddleware(c)

		// Continue to the next middleware or route handler
		c.Next()
	}
}
