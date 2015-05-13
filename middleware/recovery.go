package middleware

import (
	"github.com/gin-gonic/gin"
	h "github.com/golang-vietnam/forum/helpers"
	"log"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				errors := h.Errors{Errors: []h.Error{h.ErrInternalServer}}
				c.JSON(errors.StatusCode(), errors)
			}
		}()

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.SetErrors(c)
		c.Next()
		errors := h.GetErrors(c)
		if errors.HasError() {
			c.JSON(errors.StatusCode(), errors)
		}
	}
}
