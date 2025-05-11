package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/pkg/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
