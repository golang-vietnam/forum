package middleware

import (
	"github.com/gin-gonic/gin"
)

type Errors struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var (
	ErrInternalServer = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(ErrInternalServer.Status, ErrInternalServer)
			}
		}()

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(-1, ErrInternalServer) //will implement here
		}
	}

}
