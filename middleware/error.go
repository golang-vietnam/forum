package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if err := c.Errors.Last(); err != nil {
			if parseError, ok := err.Err.(*apiErrors.Error); ok {
				c.JSON(parseError.Status, gin.H{
					"message": parseError.Message,
					"id":      parseError.Id,
				})
				return
			}
			c.JSON(apiErrors.SERVER_ERROR.Status, gin.H{
				"message": apiErrors.SERVER_ERROR.Message,
				"id":      apiErrors.SERVER_ERROR.Id,
			})
		}
	}
}
