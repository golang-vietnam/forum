package middleware

import (
	"github.com/gin-gonic/gin"
)

type loads struct {
}

type loadsInterface interface {
	LoadUserById() gin.HandlerFunc
	LoadCategoryById() gin.HandlerFunc
}

func NewLoads() loadsInterface {
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
			return
		}

		c.Set("userData", user)
		c.Next()
	}
}

func (l *loads) LoadCategoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId := c.Param("categoryId")
		if categoryId == "" {
			panic("categoryId parameter required")
		}
		category, err := categoryResource.GetById(categoryId)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("categoryData", category)
		c.Next()
	}
}
