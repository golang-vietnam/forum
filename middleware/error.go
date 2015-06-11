package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
)

/**

	TODO:
	- Check last public error is api error
	- If has -> send a api error
	- If not -> panic to Recovery middleware
**/

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
			panic("Error is not api error")
		}
	}
}

/**

	TODO:
	- Only run on production mode
	- Log to file
	- Send SERVER_ERROR

**/

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(apiErrors.SERVER_ERROR.Status, gin.H{
					"message": apiErrors.SERVER_ERROR.Message,
					"id":      apiErrors.SERVER_ERROR.Id,
				})
				return
			}
		}()
		c.Next()
	}
}
