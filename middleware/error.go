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
			// if parseError, ok := err.Err.(*apiErrors.Error); ok {
			// 	c.JSON(parseError.Status, gin.H{
			// 		"message": parseError.Message,
			// 		"id":      parseError.Id,
			// 	})
			// 	return
			// }
			if parseError := apiErrors.ParseError(err.Err); parseError != nil {
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
	- Send ServerError

**/

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				serverError := apiErrors.ThrowError(apiErrors.ServerError)
				c.JSON(serverError.Status, gin.H{
					"message": serverError.Message,
					"id":      serverError.Id,
				})
				return
			}
		}()
		c.Next()
	}
}
