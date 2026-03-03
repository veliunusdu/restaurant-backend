package middleware

import "github.com/gin-gonic/gin"

// Authentication returns a middleware that allows all requests for now.
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
