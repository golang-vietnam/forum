package middleware

import (
	"github.com/gin-gonic/gin"
)

type loads struct {
}

type LoadsInterface interface {
	LoadUserById() gin.HandlerFunc
}

func NewLoads() LoadsInterface {
	return &loads{}
}

func (l *loads) LoadUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			panic("userId parameter required")
		}
		user, err := userResource.GetById(userId)
		if err != nil {
			c.Error(err)
			c.Abort()
		}

		c.Set("userData", user)
		c.Next()
	}
}
