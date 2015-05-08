package middleware

import (
	"github.com/gin-gonic/gin"
	_ "gopkg.in/mgo.v2"
	"log"
)

type Errors struct {
	Errors []*Error `json:"errors"`
}

func (e *Errors) StatusCode() int {
	statusCode := 500
	if len(e.Errors) > 0 {
		statusCode = e.Errors[0].Status
	}
	return statusCode
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var (
	ErrInternalServer = &Error{500, "Internal Server Error", "Something went wrong."}
	ErrBadRequest     = &Error{400, "Bad Request", "The request had bad syntax or was inherently impossible to be satisfied"}
	ErrUnauthorized   = &Error{401, "Unauthorized", "Login required"}
	ErrForbidden      = &Error{403, "Forbidden", "Access deny"}
	ErrNotFound       = &Error{404, "Not Found", "Not found anything matching the URI given"}
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				errors := &Errors{Errors: []*Error{ErrInternalServer}}
				c.JSON(errors.StatusCode(), errors)
			}
		}()

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(-1, "errors")
		}
	}

}
