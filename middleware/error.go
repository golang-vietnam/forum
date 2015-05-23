package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if err := c.Errors.Last(); err != nil {
			c.JSON(c.Writer.Status(), gin.H{
				"message": err.Error(),
			})
		}
	}
}
